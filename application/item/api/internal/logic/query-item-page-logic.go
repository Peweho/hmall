package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/item/api/internal/model"
	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"
	"hmall/application/item/api/internal/util"
	"strconv"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryItemPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryItemPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryItemPageLogic {
	return &QueryItemPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryItemPageLogic) QueryItemPage(req *types.QueryItemPageReq) (resp *types.QueryItemPageResp, err error) {
	// 0、参数处理
	var (
		sortBy   string
		isAsc    string
		page     int
		pageSize int
		total    int64
	)
	if req.SortBy == "" {
		sortBy = types.SortBy
	} else {
		sortBy = req.SortBy
	}

	if req.IsAsc == "" {
		isAsc = types.IsAsc
	} else {
		isAsc = req.IsAsc
	}

	if req.PageNo == 0 {
		page = types.Page
	} else {
		page = req.PageNo
	}

	if req.PageSize == 0 {
		pageSize = types.PageSize
	} else {
		pageSize = req.PageSize
	}

	items := make([]types.Item, 0, page)
	//1、查询缓存
	key := fmt.Sprintf("%s#%s#%s#%d#%d", types.CacheQueryItemKey, sortBy, isAsc, pageSize, page)
	keyTotal := types.CacheQueryItemTotalKey
	val, err := l.svcCtx.BizRedis.ZrangebyscoreWithScoresCtx(l.ctx, key, -time.Now().Unix(), 0)
	if len(val) > 0 && err == nil {
		_ = l.svcCtx.BizRedis.Expire(key, types.CacheQueryItemTime)
		_ = l.svcCtx.BizRedis.Expire(keyTotal, types.CacheQueryItemTime)
		//1.1、构造返回请求
		if err = rdsToItem(&items, val); err == nil {
			get, _ := l.svcCtx.BizRedis.Get(keyTotal)
			atoi, _ := strconv.Atoi(get)
			return &types.QueryItemPageResp{
				List:  items,
				Total: atoi,
				Pages: atoi / pageSize,
			}, err
		}
	} else {
		logx.Errorf("BizRedis.ZrangebyscoreWithScoresAndLimitCtx: %v, error: %v", key, err)
	}

	// 2、查询数据库
	//查询商品
	res, err := l.svcCtx.ItemModel.QueryItemPage(l.ctx, page, pageSize, sortBy, isAsc)
	if err != nil {
		logx.Errorf("ItemModel.QueryItemPage, error: %v", err)
		return nil, err
	}
	//查询商品总数量
	err = l.svcCtx.ItemModel.GetItemTotal(l.ctx, &total)
	if err != nil {
		logx.Errorf("ItemModel.GetItemTotal, error: %v", err)
		return nil, err
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	threading.GoSafe(func() {
		defer wg.Done()
		//3、写缓存
		err = l.WriteCache(key, &res, keyTotal, total)
		if err != nil {
			panic(err)
		}
	})

	threading.GoSafe(func() {
		defer wg.Done()
		//4、构造数据
		for _, val := range res {
			items = append(items, util.ItemDTO_To_Item(val))
		}
	})

	wg.Wait()
	return &types.QueryItemPageResp{
		List:  items,
		Pages: int(total) / pageSize,
		Total: int(total),
	}, nil
}

// 将redis pair转为返回对象
func rdsToItem(items *[]types.Item, pair []redis.Pair) error {

	for _, val := range pair {
		var item model.ItemDTO
		err := json.Unmarshal([]byte(val.Key), &item)
		if err != nil {
			logx.Errorf("json.Unmarshal: %v, error: %v", val.Key, err)
			return err
		}
		*items = append(*items, util.ItemDTO_To_Item(item))
	}
	return nil
}

// 写缓存
func (l *QueryItemPageLogic) WriteCache(key string, data *[]model.ItemDTO, keyTotal string, total int64) error {
	pair := make([]redis.Pair, 0, len(*data))
	for _, val := range *data {
		marshal, err := json.Marshal(val)
		if err != nil {
			logx.Errorf("json.Marshal, error: %v", err)
			return err
		}
		pair = append(pair, redis.Pair{
			Key:   string(marshal),
			Score: -val.Price,
		})
	}
	//将查询到的数据写入缓存
	_, err := l.svcCtx.BizRedis.ZaddsCtx(l.ctx, key, pair...)
	if err != nil {
		logx.Errorf("BizRedis.ZaddsCtx: %v, error: %v", key, err)
		return err
	}
	//记录数据总量
	_ = l.svcCtx.BizRedis.Set(keyTotal, strconv.FormatInt(total, 10))
	_ = l.svcCtx.BizRedis.Expire(key, types.CacheQueryItemTime)
	_ = l.svcCtx.BizRedis.Expire(keyTotal, types.CacheQueryItemTime)

	return nil
}
