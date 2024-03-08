package logic

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/item/rpc/internal/model"
	"hmall/application/item/rpc/internal/svc"
	"hmall/application/item/rpc/pb"
	"hmall/application/item/rpc/types"
	"hmall/pkg/util"
	"hmall/pkg/xcode"
	"log"
	"strconv"
	"sync"
	"time"
)

type FindItemByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	Bloom   *bloom.Filter
	KqCache *PusherCacheLogic
}

func NewFindItemByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindItemByIdsLogic {
	return &FindItemByIdsLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		Bloom:   bloom.New(svcCtx.BizRedis, types.ItemBloomKey, 20*svcCtx.Config.ItemNums),
		KqCache: NewPusherCacheLogic(ctx, svcCtx),
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
	for _, v := range in.Ids {
		//布隆过滤器先过滤
		exists, _ := l.Bloom.ExistsCtx(l.ctx, []byte(v))
		if !exists {
			continue
		}

		wg.Add(1)
		threading.GoSafe(func() {
			defer wg.Done()
			err := l.ReadCacheV2(&items, v)
			log.Println(items)
			if err != nil {
				panic(err)
			}
		})
	}
	wg.Wait()

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

// 读取一个缓存
func (l *FindItemByIdsLogic) ReadCacheV2(items *[]*pb.Items, id string) error {
	for {
		key := util.CacheKey(types.CacheItemKey, id)
		cacheItem, _ := l.svcCtx.BizRedis.Hgetall(key)
		_, ok := cacheItem[types.CacheItemFields]
		if ok {
			//如果数据不为空，那么立即返回结果，并异步执行"取数据"
			//取数据,mq实现
			_ = l.KqCache.PusherCache(id)
			//获取结果
			l.getResult(items, cacheItem)
			return nil
		} else {
			lockUtils, ok := cacheItem[types.CacheItemLockUils]
			var (
				now  int64
				lock int64
			)
			if ok {
				now = time.Now().UnixMilli()
				lock, _ = strconv.ParseInt(lockUtils, 10, 64)
			}

			//没有锁、锁为过期标志、小于当前时间表示锁失效
			if !ok || lockUtils == types.CacheItemDeadLine || lock < now {
				err := l.updateCache(id, items)
				return err
			} else {
				//如果数据为空，且被锁定，则睡眠100ms后，重新查询
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

// 取数据并获得结果
func (l *FindItemByIdsLogic) updateCache(id string, items *[]*pb.Items) error {
	key := util.CacheKey(types.CacheItemKey, id)
	Uuid := uuid.New().String()
	val := map[string]string{
		types.CacheItemLockUils: strconv.FormatInt(time.Now().UnixMilli()+int64(types.CacheItemLockUilsTime), 10),
		types.CacheItemOwner:    Uuid,
	}
	err := l.svcCtx.BizRedis.Hmset(key, val)
	if err != nil {
		logx.Errorf("BizRedis.Hmset: key=%v,value=%v,error: %v", key, val, err)
		return err
	}

	//数据库查询
	res, err := l.svcCtx.ItemModel.FindItemById(l.ctx, id)
	if err != nil {
		logx.Errorf("ItemModel.FindItemById: %v ,error: %v", id, err)
		return err
	}

	*items = append(*items, ItemDTO_To_Item(res))

	//写缓存
	marshal, err := json.Marshal(res)
	if err != nil {
		logx.Errorf("json.Marshal: %v , error: ,%v", res, err)
		return err
	}

	cacheItem, _ := l.svcCtx.BizRedis.Hgetall(key)
	if cacheItem[types.CacheItemOwner] == Uuid {
		set := map[string]string{
			types.CacheItemFields: string(marshal),
			types.CacheItemStatus: strconv.FormatInt(res.Status, 10),
			types.CacheItemStock:  strconv.FormatInt(res.Stock, 10),
		}
		if err := l.svcCtx.BizRedis.Hmset(key, set); err != nil {
			logx.Errorf("BizRedis.Hmset: key=%v,value=%v, error: %v", key, set, err)
			return nil
		}
	}
	return nil
}

// 获取结果
func (l *FindItemByIdsLogic) getResult(items *[]*pb.Items, cacheItem map[string]string) {
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
	temp.Stock = stock

	//设置状态
	status, err := strconv.ParseInt(cacheItem[types.CacheItemStatus], 10, 64)
	if err != nil {
		logx.Errorf("strconv.ParseInt: %v , error: ,%v", cacheItem[types.CacheItemStatus], err)
		panic(err)
	}
	temp.Status = status
	*items = append(*items, &temp)
}
