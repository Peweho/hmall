package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/item/rpc/internal/model"
	"hmall/application/item/rpc/internal/svc"
	"hmall/application/item/rpc/pb"
	"hmall/application/item/rpc/types"
	"hmall/pkg/util"
	"hmall/pkg/xcode"
	"strconv"
	"sync"
)

type FindItemByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	Bloom *bloom.Filter
}

func NewFindItemByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindItemByIdsLogic {
	return &FindItemByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		Bloom:  bloom.New(svcCtx.BizRedis, types.ItemBloomKey, 20*svcCtx.Config.ItemNums),
	}
}

func (l *FindItemByIdsLogic) FindItemByIds(in *pb.FindItemByIdsReq) (*pb.FindItemByIdsResp, error) {
	// 1、校验参数
	if len(in.Ids) == 0 {
		return nil, xcode.New(200, "ids为空")
	}
	items := make([]*pb.Items, 0, len(in.Ids))

	//2、查询缓存 并且 构造新的请求ids
	wg := &sync.WaitGroup{}
	newIds, err := l.ReadCache(&items, in.Ids, wg)
	if err != nil {
		return nil, err
	}

	wg.Wait()
	if len(newIds) == 0 {
		return &pb.FindItemByIdsResp{Data: items}, nil
	}

	//2、查询数据
	res, err := l.svcCtx.ItemModel.FindItemByIds(l.ctx, in.Ids)
	if err != nil {
		logx.Errorf("ItemModel.FindItemByIds: %v ,error: %v", in.Ids, err)
		return nil, err
	}

	//3、写缓存
	threading.NewWorkerGroup(func() {
		_ = l.WriteCache(res)
	}, 1).Start()

	//4、类型转换
	for _, v := range res {
		items = append(items, ItemDTO_To_Item(v))
	}

	//5、返回响应
	return &pb.FindItemByIdsResp{Data: items}, nil
}

func ItemDTO_To_Item(item model.ItemDTO) *pb.Items {
	return &pb.Items{
		Id:           item.Id,
		Brand:        item.Brand,
		Category:     item.Category,
		CommentCount: item.CommentCount,
		Image:        item.Image,
		IsAD:         item.IsAD,
		Name:         item.Name,
		Price:        item.Price,
		Sold:         item.Sold,
		Spec:         item.Spec,
		Stock:        item.Stock,
	}
}

// 构造对应的缓存键
func CacheIds(id string) string {
	return fmt.Sprintf("%s#%s", types.CacheItemKey, id)
}

// 读缓存
func (l *FindItemByIdsLogic) ReadCache(items *[]*pb.Items, ids []string, wg *sync.WaitGroup) ([]string, error) {
	newIds := make([]string, 0, len(ids))
	for _, v := range ids {
		key := util.CacheKey(types.CacheItemKey, v)

		//使用布隆过滤器判断id是否存在于数据库
		bloomOk, _ := l.Bloom.ExistsCtx(l.ctx, []byte(v))
		if !bloomOk {
			continue
		}

		ok, _ := l.svcCtx.BizRedis.Exists(key)
		if !ok {
			//缓存不存在，将对应id存储到newIds中
			newIds = append(newIds, v)
			continue
		}
		wg.Add(1)
		threading.GoSafe(func() {
			defer wg.Done()
			//判对应商品是存在，存在加入items
			_ = l.svcCtx.BizRedis.Expire(key, types.CacheItemTime) //设置有效期
			cacheItem, _ := l.svcCtx.BizRedis.Hgetall(key)
			var temp pb.Items
			err := json.Unmarshal([]byte(cacheItem[types.CacheItemFields]), &temp)
			if err != nil {
				logx.Errorf("json.Unmarshal: %v , error: ,%v", cacheItem, err)
				panic(err)
			}
			//设置库存
			stock, err := strconv.ParseInt(cacheItem[types.CacheItemStock], 10, 64)
			if err != nil {
				logx.Errorf("strconv.ParseInt: %v , error: ,%v", cacheItem[types.CacheItemStock], err)
				panic(err)
			}
			temp.Stock = int64(stock)

			//设置状态
			status, err := strconv.ParseInt(cacheItem[types.CacheItemStatus], 10, 64)
			if err != nil {
				logx.Errorf("strconv.ParseInt: %v , error: ,%v", cacheItem[types.CacheItemStatus], err)
				panic(err)
			}
			temp.Status = status
			*items = append(*items, &temp)
		})
	}
	return newIds, nil
}

// 写缓存
func (l *FindItemByIdsLogic) WriteCache(items []model.ItemDTO) error {
	for _, v := range items {
		marshal, err := json.Marshal(v)
		if err != nil {
			logx.Errorf("json.Marshal: %v , error: ,%v", v, err)
			return err
		}

		key := util.CacheKey(types.CacheItemKey, strconv.FormatInt(v.Id, 10))
		if err := l.svcCtx.BizRedis.Hmset(key, map[string]string{
			types.CacheItemFields: string(marshal),
			types.CacheItemStatus: strconv.FormatInt(v.Status, 10),
			types.CacheItemStock:  strconv.FormatInt(v.Stock, 10),
		}); err != nil {
			logx.Errorf("BizRedis.Hmset: %v, error: %v", key, err)
		}

		//设置有效期
		_ = l.svcCtx.BizRedis.Expire(key, types.CacheItemTime)
	}

	return nil
}
