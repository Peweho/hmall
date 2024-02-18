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
		sort.Sort(model.CartPOArr(catr))
		wg.Done()
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
			Image:      catr[i].Image,
			ItemId:     itemId,
			Name:       val.Name,
			Price:      catr[i].Price,
			NewPrice:   int(val.Price),
			Num:        catr[i].Num,
			Spec:       val.Spec,
			CreateTime: time.Now().Format(time.DateTime),
			Status:     int(val.Status),
			Stock:      int(val.Stock),
		}
		//3、写缓存
		threading.GoSafe(func() {
			wg.Add(1)
			marshal, err := json.Marshal(temp)
			if err != nil {
				logx.Errorf("json.Marshal: %v, error: %v", temp, err)
				return
			}
			_, err = l.svcCtx.BizRedis.Zadd(key, time.Now().Unix(), string(marshal))
			if err != nil {
				logx.Errorf("BizRedis.Zadd: %v, error: %v", string(marshal), err)
				return
			}
			wg.Done()
		})
		CartItems = append(CartItems, temp)
	}

	//4、返回响应
	wg.Wait()
	return &types.QueryCartResp{Items: CartItems}, nil
}

// 查询缓存
func (l *QueryCartLogic) QueryCache(key string) (resp *types.QueryCartResp, err error) {
	num, _ := l.svcCtx.BizRedis.Zcard(key)
	CartItems := make([]types.ItemDTO, 0, num)
	redisPair, err := l.svcCtx.BizRedis.ZrangebyscoreWithScores(key, 0, time.Now().Unix())
	if err != nil {
		logx.Errorf(":%v, error: %v", key, err)
		return nil, err
	}
	for _, val := range redisPair {
		CartItem := types.ItemDTO{}
		if err := json.Unmarshal([]byte(val.Key), &CartItem); err != nil {
			logx.Errorf("json.Unmarshal:%v, error: %v", val.Key, err)
			return nil, err
		}
		CartItems = append(CartItems, CartItem)
	}

	if err = l.svcCtx.BizRedis.Expire(key, types.CacheCartTime); err != nil {
		logx.Errorf("BizRedis.Expire:%v, error: %v", key, err)
		return nil, err
	}
	return &types.QueryCartResp{
		Items: CartItems,
	}, nil
}
