package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hmall/application/pay/rpc/internal/utils"
	"time"

	"hmall/application/pay/rpc/internal/svc"
	"hmall/application/pay/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePayOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePayOrderLogic {
	return &UpdatePayOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePayOrderLogic) UpdatePayOrder(in *pb.UpdatePayOrderReq) (*pb.UpdatePayOrderResp, error) {
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		update := map[string]any{
			"status":           utils.PaySuccess,
			"pay_success_time": time.Now(),
			"pay_type":         in.PayType,
		}
		if err2 := l.svcCtx.PayModel.UpdatePayOrder(l.ctx, update); err2 != nil {
			logx.Errorf("PayModel.UpdatePayOrder: %v,error : %v", update, err2)
			return err2
		}
		return nil
	}); err != nil {
		if err == dtmcli.ErrFailure {
			return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.UpdatePayOrderResp{}, nil
}
