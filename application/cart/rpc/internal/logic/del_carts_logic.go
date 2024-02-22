package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"hmall/application/cart/rpc/internal/svc"
	"hmall/application/cart/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCartsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCartsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCartsLogic {
	return &DelCartsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCartsLogic) DelCarts(in *pb.DelCartsReq) (*pb.DelCartsResp, error) {
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()

	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		if err = l.svcCtx.CartModel.DelCartsByUidItemId(l.ctx, in.Usr, in.ItemId); err != nil {
			logx.Errorf("CartModel.DelCartsByIds: %v, error: %v", in.Usr, err)
			return err
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}

	return &pb.DelCartsResp{}, nil
}
