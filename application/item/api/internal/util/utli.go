package util

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/item/api/internal/model"
	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"
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

func ItemDTO_To_Item(item model.ItemDTO) types.Item {
	return types.Item{
		Id:           int(item.Id),
		Brand:        item.Brand,
		Category:     item.Category,
		CommentCount: int(item.CommentCount),
		Image:        item.Image,
		IsAD:         item.IsAD,
		Name:         item.Name,
		Price:        int(item.Price),
		Sold:         int(item.Sold),
		Spec:         item.Spec,
		Status:       int(item.Status),
		Stock:        int(item.Stock),
	}
}

// ItemReqAndResp 转 ItemDTO
func ItemReqAndResp_To_ItemDTO(req *types.ItemReqAndResp) model.ItemDTO {
	return model.ItemDTO{
		Id:           int64(req.Id),
		Brand:        req.Brand,
		Category:     req.Category,
		CommentCount: int64(req.CommentCount),
		Image:        req.Image,
		IsAD:         req.IsAD,
		Name:         req.Name,
		Price:        int64(req.Price),
		Sold:         int64(req.Sold),
		Spec:         req.Spec,
		Status:       int64(req.Status),
		Stock:        int64(req.Stock),
	}
}
