package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/item/rpc/internal/model"
	"hmall/application/item/rpc/internal/svc"
	"hmall/application/item/rpc/types"
	"hmall/application/item/rpc/types/service"
	"hmall/pkg/xcode"
	"strconv"
)

type FindItemByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindItemByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindItemByIdsLogic {
	return &FindItemByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindItemByIdsLogic) FindItemByIds(in *service.FindItemByIdsReq) (*service.FindItemByIdsResp, error) {
	// 1、校验参数
	if len(in.Ids) == 0 {
		return nil, xcode.New(200, "ids为空")
	}

	items := make([]*service.Items, 0, len(in.Ids))

	//2、查询缓存 并且 构造新的请求ids
	newIds, err := l.ReadCache(&items, in.Ids)
	if err != nil {
		return nil, err
	}

	if len(newIds) == 0 {
		return &service.FindItemByIdsResp{Data: items}, nil
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
	}, 1)

	//4、类型转换
	for _, v := range res {
		items = append(items, ItemDTO_To_Item(v))
	}

	//5、返回响应
	return &service.FindItemByIdsResp{Data: items}, nil
}

func ItemDTO_To_Item(item model.ItemDTO) *service.Items {
	return &service.Items{
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
		Status:       item.Status,
		Stock:        item.Stock,
	}
}

// 构造对应的缓存键
func CacheIds(id string) string {
	return fmt.Sprintf("%s#%s", types.CacheItemKey, id)
}

// 读缓存
func (l *FindItemByIdsLogic) ReadCache(items *[]*service.Items, ids []string) ([]string, error) {

	newIds := make([]string, 0, len(ids))
	for _, v := range ids {
		key := CacheIds(v)
		ok, _ := l.svcCtx.BizRedis.Exists(key)
		if !ok {
			//缓存不存在，将对应id存储到newIds中
			newIds = append(newIds, v)
			continue
		}

		//判对应商品是存在，存在加入items
		_ = l.svcCtx.BizRedis.Expire(key, types.CacheItemTime) //设置有效期
		cacheItem, _ := l.svcCtx.BizRedis.Get(key)
		var temp service.Items
		err := json.Unmarshal([]byte(cacheItem), &temp)
		if err != nil {
			logx.Errorf("json.Unmarshal: %v , error: ,%v", cacheItem, err)
			return nil, err
		}
		*items = append(*items, &temp)
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
		key := CacheIds(strconv.Itoa(int(v.Id)))
		err = l.svcCtx.BizRedis.Set(key, string(marshal))
		if err != nil {
			logx.Errorf("BizRedis.Set: %v , error: ,%v", string(marshal), err)
			return err
		}
		_ = l.svcCtx.BizRedis.Expire(key, types.CacheItemTime) //设置有效期
	}
	return nil
}
