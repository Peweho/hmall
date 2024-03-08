package mqs

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/item/mq/internal/svc"
	"hmall/application/item/mq/internal/types"
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

// Kq补偿类型
const (
	KqCacheAll = iota
	KqCachePart
	KqCacheStock
	KqCacheStatus
	KqCacheDel
)

type KqCacheMsg struct {
	Code   int //补偿类型：0、全字段添加；1、更新部分字段；2，更新库存；3、更新状态;4、删除
	Field  string
	Stock  string
	Status string
	Key    string
}

// 更新缓存
func (l *PaymentSuccess) Consume(_, data string) error {
	msg := &KqCacheMsg{}
	if err := json.Unmarshal([]byte(data), msg); err != nil {
		logx.Errorf(": %v, error； %v", data, err)
		return err
	}
	//获取控制字段
	cacheItem, _ := l.svcCtx.BizRedis.Hgetall(msg.Key)
	_, ok := cacheItem[types.CacheItemFields]
	lockUtils, ok := cacheItem[types.CacheItemLockUils]

	//判断是否过期
	if !ok || lockUtils != types.CacheItemDeadLine {
		return nil
	}

	//写入控制信息
	Uuid := uuid.New().String()
	val := map[string]string{
		types.CacheItemLockUils: strconv.FormatInt(time.Now().UnixMilli()+int64(types.CacheItemLockUilsTime), 10),
		types.CacheItemOwner:    Uuid,
	}
	err := l.svcCtx.BizRedis.Hmset(msg.Key, val)
	if err != nil {
		logx.Errorf("BizRedis.Hmset: key=%v,value=%v,error: %v", msg.Key, val, err)
		return err
	}

	//进行更新
	switch msg.Code {
	case KqCacheAll:
		if err := l.CacheAll(msg); err != nil {
			logx.Errorf("l.CacheAll: %v, error: %v", *msg, err)
			return err
		}
	case KqCachePart:
		if err := l.CacheField(msg); err != nil {
			logx.Errorf("l.CacheField: %v, error: %v", *msg, err)
			return err
		}
	case KqCacheStock:
		if err := l.CacheStock(msg); err != nil {
			logx.Errorf("l.CacheStoc: %v, error: %v", *msg, err)
			return err
		}
	case KqCacheStatus:
		if err := l.CacheStatus(msg); err != nil {
			logx.Errorf("l.CacheStatus: %v, error: %v", *msg, err)
			return err
		}
	case KqCacheDel:
		if err := l.CacheDel(msg); err != nil {
			logx.Errorf("l.CacheDel: %v, error: %v", *msg, err)
			return err
		}
	default:
		panic("传输code不正确")
	}

	return nil
}

// 全字段添加
func (l *PaymentSuccess) CacheAll(msg *KqCacheMsg) error {
	if err := l.svcCtx.BizRedis.Hmset(msg.Key, map[string]string{
		types.CacheItemFields: msg.Field,
		types.CacheItemStatus: msg.Status,
		types.CacheItemStock:  msg.Stock,
	}); err != nil {
		logx.Errorf("BizRedis.Hmset: %v, error: %v", msg.Key, err)
		return err
	}
	return nil
}

// 更新部分字段
func (l *PaymentSuccess) CacheField(msg *KqCacheMsg) error {
	if err := l.svcCtx.BizRedis.Hset(msg.Key, types.CacheItemFields, msg.Field); err != nil {
		logx.Errorf("BizRedis.Hset: %v, error: %v", msg.Key, err)
		return err
	}
	return nil
}

// 更新库存
func (l *PaymentSuccess) CacheStock(msg *KqCacheMsg) error {
	if err := l.svcCtx.BizRedis.Hset(msg.Key, types.CacheItemStock, msg.Stock); err != nil {
		logx.Errorf("BizRedis.Hset: %v, error: %v", msg.Key, err)
		return err
	}
	return nil
}

// 更新状态
func (l *PaymentSuccess) CacheStatus(msg *KqCacheMsg) error {
	if err := l.svcCtx.BizRedis.Hset(msg.Key, types.CacheItemStatus, msg.Status); err != nil {
		logx.Errorf("BizRedis.Hset: %v, error: %v", msg.Key, err)
		return err
	}
	return nil
}

// 删除
func (l *PaymentSuccess) CacheDel(msg *KqCacheMsg) error {
	if _, err := l.svcCtx.BizRedis.Del(msg.Key); err != nil {
		logx.Errorf("BizRedis.Del: %v, error: %v", msg.Key, err)
		return err
	}
	return nil
}
