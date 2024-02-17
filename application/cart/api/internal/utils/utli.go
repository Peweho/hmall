package utils

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/cart/api/internal/svc"
	"hmall/application/cart/api/internal/types"
	"hmall/application/item/rpc/item"
	"log"
	"strconv"
)

//提供公共方法

func UpdateCache(ctx context.Context, svcCtx *svc.ServiceContext, id int) error {
	strId := strconv.Itoa(id)
	key := CacheIds(strId)
	//1、删除缓存
	_, err := svcCtx.BizRedis.Del(key)
	log.Println("删除缓存: %v", key)
	//2、添加缓存
	_, err = svcCtx.ItemRPC.FindItemByIds(ctx, &item.FindItemByIdsReq{Ids: []string{strId}})
	if err != nil {
		logx.Errorf("ItemRPC.FindItemByIds: %v, error: %v", strId, err)
		return err
	}
	return nil
}

// 构造对应的缓存键
func CacheIds(id string) string {
	return fmt.Sprintf("%s#%s", types.CacheItemKey, id)
}
