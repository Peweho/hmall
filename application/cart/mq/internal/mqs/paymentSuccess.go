package mqs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/cart/mq/internal/model"
	"hmall/application/cart/mq/internal/svc"
	"hmall/application/cart/mq/internal/types"
	"hmall/application/item/rpc/item"
	"strconv"
	"time"
)

type PaymentSuccess struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentSuccess(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentSuccess {
	return &PaymentSuccess{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type KqMsg struct {
	Category string // 1,删缓存；0，加缓存
	Data     *model.CartPO
}

type CartItem struct {
	CreateTime string `json:"create_time"`
	Id         int    `json:"id"`
	Image      string `json:"image"`
	ItemId     int    `json:"itemId"`
	Name       string `json:"name"`
	NewPrice   int    `json:"newPrice"`
	Num        int    `json:"num"`
	Price      int    `json:"price"`
	Spec       string `json:"spec"`
	Status     int    `json:"status"`
	Stock      int    `json:"stock"`
}

// 更新缓存
func (l *PaymentSuccess) Consume(_, str string) error {
	// 1、序列化数据
	msg := &KqMsg{}
	err := json.Unmarshal([]byte(str), msg)
	if err != nil {
		logx.Errorf("json.Unmarshal: %v, error: %v", str, err)
		return err
	}
	key := CacheIds(strconv.Itoa(msg.Data.UserId))
	// 2、分类
	switch msg.Category {
	case types.MSgAddCache:
		//查询商品数据
		itemData, err := l.svcCtx.ItemRPC.FindItemByIds(l.ctx, &item.FindItemByIdsReq{Ids: []string{strconv.Itoa(msg.Data.ItemId)}})
		if err != nil {
			logx.Errorf("ItemRPC.FindItemByIds: %v, error: %v", msg.Data.ItemId, err)
			return err
		}
		// 如果商品已经不存在就忽略
		if len(itemData.Data) == 0 {
			return nil
		}

		//构造购物车商品数据
		CartItem := &CartItem{
			Id:         msg.Data.Id,
			Num:        msg.Data.Num,
			ItemId:     msg.Data.ItemId,
			Image:      itemData.Data[0].Image,
			CreateTime: time.Now().Format("2006-01-02 15:04:05"),
			Spec:       itemData.Data[0].Spec,
			NewPrice:   int(itemData.Data[0].Price),
			Status:     int(itemData.Data[0].Status),
			Stock:      int(itemData.Data[0].Stock),
			Name:       itemData.Data[0].Name,
		}

		//序列化
		marshal, err := json.Marshal(CartItem)
		if err != nil {
			logx.Errorf("json.Marshal: %v, error: %v", *msg, err)
			return err
		}
		//增加缓存
		if err = l.AddCache(key, string(marshal), strconv.Itoa(CartItem.Id)); err != nil {
			logx.Errorf("AddCache: %v, error: %v", *msg, err)
			return err
		}
	case types.MSgDelCache:
		if err = l.DelCache(key, strconv.Itoa(msg.Data.Id)); err != nil {
			logx.Errorf("DelCache: %v, error: %v", *msg, err)
			return err
		}
	}
	return nil
}

// 添加缓存
func (l *PaymentSuccess) AddCache(key, data string, id string) error {
	err := l.svcCtx.BizRedis.Hset(key, id, data)
	if err != nil {
		logx.Errorf("BizRedis.Hset: %v, error: %v", key, err)
		return err
	}
	if err = l.svcCtx.BizRedis.Expire(key, types.CacheCartTime); err != nil {
		logx.Errorf("BizRedis.Expire:%v, error: %v", key, err)
		return err
	}
	return nil
}

// 删除缓存
func (l *PaymentSuccess) DelCache(key, id string) error {
	_, err := l.svcCtx.BizRedis.Hdel(key, id)
	if err != nil {
		logx.Errorf("BizRedis.Zadd: %v, error: %v", key, err)
		return err
	}
	return nil
}

func CacheIds(id string) string {
	return fmt.Sprintf("%s#%s", types.CacheCartKey, id)
}
