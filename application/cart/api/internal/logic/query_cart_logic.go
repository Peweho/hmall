package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/cart/api/internal/model"
	"hmall/application/cart/api/internal/utils"
	"hmall/application/item/rpc/item"
	"hmall/pkg/util"
	"hmall/pkg/xcode"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/cart/api/internal/svc"
	"hmall/application/cart/api/internal/types"
	ItemUtils "hmall/application/item/rpc/utils"
)

type QueryCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCartLogic {
	return &QueryCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryCartLogic) QueryCart() (resp *types.QueryCartResp, err error) {
	// 1、查询缓存
	usr, err := util.GetUsr(l.ctx, types.JwtKey)
	if err != nil {
		logx.Errorf("util.GetUsr, error: %v", err)
		return nil, xcode.New(types.Unauthorized, "")
	}
	key := utils.CacheIds(strconv.Itoa(usr))
	exists, _ := l.svcCtx.BizRedis.Exists(key)
	if exists {
		return l.QueryCache(key)
	}

	//2、查询数据库
	//查询购物车
	catr, err := l.svcCtx.CartModel.QueryCatr(l.ctx, usr)
	if err != nil {
		logx.Errorf("CartModel.QueryCatr:%v, error: %v", usr, err)
		return nil, err
	}
	//构建商品id集合
	ids := make([]string, 0, len(catr))
	for _, val := range catr {
		ids = append(ids, strconv.Itoa(val.ItemId))
	}
	//对cart排序根据itemid
	wg := sync.WaitGroup{}
	wg.Add(1)
	threading.GoSafe(func() {
		defer wg.Done()
		sort.Sort(model.CartPOArr(catr))
	})

	//查找商品
	itemRpcresp, err := l.svcCtx.ItemRPC.FindItemByIds(l.ctx, &item.FindItemByIdsReq{Ids: ids})
	if err != nil {
		logx.Errorf("ItemRPC.FindItemByIds: %v, error: %v", ids, err)
		return nil, xcode.New(types.NotFound, "")
	}
	//对商品集合排序，保证与购物车一致
	sort.Sort(ItemUtils.ItemsArr(itemRpcresp.Data))
	wg.Wait()
	chItem := make(chan *types.ItemDTO, len(catr))
	//构造返回数据
	CartItems := make([]types.ItemDTO, 0, len(catr))
	for i, val := range itemRpcresp.Data {
		itemId := int(val.Id)
		if itemId != catr[i].ItemId { //检查商品ID是否一致
			logx.Errorf("error: %v", "购物车中的商品id与rpc查到商品id不一致")
			return nil, xcode.New(types.NotFound, "")
		}
		temp := types.ItemDTO{
			Id:         catr[i].Id,
			Image:      val.Image,
			ItemId:     itemId,
			Name:       val.Name,
			Price:      catr[i].Price,
			NewPrice:   int(val.Price),
			Num:        catr[i].Num,
			Spec:       val.Spec,
			CreateTime: catr[i].CreatedAt.Format(time.DateTime),
			Status:     int(val.Status),
			Stock:      int(val.Stock),
		}
		chItem <- &temp
		//3、写缓存
		wg.Add(1)
		threading.GoSafe(func() {
			defer wg.Done()
			date := <-chItem
			marshal, err := json.Marshal(date)
			if err != nil {
				logx.Errorf("json.Marshal: %v, error: %v", date, err)
				return
			}

			if err = l.svcCtx.BizRedis.Hset(key, strconv.Itoa(date.Id), string(marshal)); err != nil {
				logx.Errorf("BizRedis.Hset: %v, error: %v", string(marshal), err)
				//调用mq服务删除缓存
				msg := &utils.KqMsg{
					Category: types.MSgAddCompleteCache,
					Data:     &model.CartPO{Id: date.Id, UserId: usr},
					Else:     string(marshal),
				}
				pusher := utils.NewPusherLogic(l.ctx, l.svcCtx)
				if err1 := pusher.UpdateCache(msg); err1 != nil {
					logx.Errorf("pusher.UpdateCache: %v, error: %v", *msg, err1)
					return
				}
			}
		})
		CartItems = append(CartItems, temp)
	}
	close(chItem)
	//4、返回响应
	wg.Wait()
	_ = l.svcCtx.BizRedis.Expire(key, types.CacheCartTime)

	return &types.QueryCartResp{Items: CartItems}, nil
}

// 查询缓存
func (l *QueryCartLogic) QueryCache(key string) (resp *types.QueryCartResp, err error) {
	redisData, err := l.svcCtx.BizRedis.Hgetall(key)
	if err != nil {
		logx.Errorf(":%v, error: %v", key, err)
		return nil, err
	}
	num := len(redisData)
	CartItems := make([]types.ItemDTO, 0, num)
	for _, val := range redisData {
		CartItem := types.ItemDTO{}
		if err := json.Unmarshal([]byte(val), &CartItem); err != nil {
			logx.Errorf("json.Unmarshal:%v, error: %v", val, err)
			return nil, err
		}
		CartItems = append(CartItems, CartItem)
	}

	_ = l.svcCtx.BizRedis.Expire(key, types.CacheCartTime)

	return &types.QueryCartResp{
		Items: CartItems,
	}, nil
}
