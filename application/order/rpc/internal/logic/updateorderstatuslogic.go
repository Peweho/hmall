package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hmall/application/order/rpc/internal/svc"
	"hmall/application/order/rpc/internal/utils"
	"hmall/application/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusLogic {
	return &UpdateOrderStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderStatusLogic) UpdateOrderStatus(in *pb.UpdateOrderStatusReq) (*pb.UpdateOrderStatusResp, error) {
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		if err2 := l.svcCtx.OrderModel.UpdateOrderStatusById(l.ctx, in.Id, utils.Paied); err2 != nil {
			logx.Errorf("OrderModel.UpdateOrderStatusById: %v , error: %v", in.Id, err2)
			return err2
		}
		return nil
	}); err != nil {
		if err == dtmcli.ErrFailure {
			return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.UpdateOrderStatusResp{}, nil
}
