package logic

import (
	"context"
	"encoding/json"
	"hmall/application/order/rpc/internal/model"
	"hmall/application/order/rpc/internal/utils"
	"time"

	"hmall/application/order/rpc/internal/svc"
	"hmall/application/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOrderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOrderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOrderByIdLogic {
	return &FindOrderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindOrderByIdLogic) FindOrderById(in *pb.FindOrderByIdReq) (*pb.FindOrderByIdResp, error) {
	//查缓存
	key := utils.CacheKey(int(in.Id))
	exists, _ := l.svcCtx.BizRedis.Exists(key)
	if exists {
		resp, err := l.FindCacheOrder(key)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	//查数据库
	order, err := l.svcCtx.OrderModel.FindOrderById(l.ctx, in.Id)
	if err != nil {
		logx.Errorf(": %v, error: %v", in.Id, err)
		return &pb.FindOrderByIdResp{}, err
	}

	//写缓存
	err = l.WriteCacheOrder(key, &order)
	if err != nil {
		return &pb.FindOrderByIdResp{}, err
	}

	return &pb.FindOrderByIdResp{
		Id:          order.Id,
		PayTime:     order.Pay_time.Format(time.DateTime),
		PaymentType: int64(order.Payment_type),
		Status:      int64(order.Status),
		TotalFee:    int64(order.Total_fee),
		UserId:      order.User_id,
		CloseTime:   order.Close_time.Format(time.DateTime),
		CommentTime: order.Comment_time.Format(time.DateTime),
		ConsignTime: order.Consign_time.Format(time.DateTime),
		CreateTime:  order.CreatedAt.Format(time.DateTime),
		EndTime:     order.End_time.Format(time.DateTime),
	}, nil
}

func (l *FindOrderByIdLogic) FindCacheOrder(key string) (*pb.FindOrderByIdResp, error) {
	get, err := l.svcCtx.BizRedis.Get(key)
	if err != nil {
		logx.Errorf("BizRedis.Get: %v, error: %v", key, err)
		return nil, err
	}
	_ = l.svcCtx.BizRedis.Expire(key, utils.CacheOrderTime)
	var res model.OrderPO
	if err = json.Unmarshal([]byte(get), &res); err != nil {
		logx.Errorf("json.Unmarshal: %v, error: %v", get, err)
		return nil, err
	}
	return &pb.FindOrderByIdResp{
		Id:          res.Id,
		PayTime:     res.Pay_time.Format(time.DateTime),
		PaymentType: int64(res.Payment_type),
		Status:      int64(res.Status),
		TotalFee:    int64(res.Total_fee),
		UserId:      res.User_id,
		CloseTime:   res.Close_time.Format(time.DateTime),
		CommentTime: res.Comment_time.Format(time.DateTime),
		ConsignTime: res.Consign_time.Format(time.DateTime),
		CreateTime:  res.CreatedAt.Format(time.DateTime),
		EndTime:     res.End_time.Format(time.DateTime),
	}, nil
}

func (l *FindOrderByIdLogic) WriteCacheOrder(key string, order *model.OrderPO) error {
	marshal, err := json.Marshal(order)
	if err != nil {
		logx.Errorf("json.Marshal: %v, error: %v", order, err)
		return err
	}
	if err = l.svcCtx.BizRedis.Set(key, string(marshal)); err != nil {
		logx.Errorf("BizRedis.Set: %v, error :%v", string(marshal), err)
		return err
	}
	_ = l.svcCtx.BizRedis.Expire(key, utils.CacheOrderTime)
	return nil
}
