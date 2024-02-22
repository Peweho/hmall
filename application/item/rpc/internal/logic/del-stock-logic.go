package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hmall/application/item/rpc/internal/svc"
	"hmall/application/item/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelStockLogic {
	return &DelStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelStockLogic) DelStock(in *service.DelStockReq) (*service.DelStockResp, error) {
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()

	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		for _, val := range in.Detail {
			item, err := l.svcCtx.ItemModel.DecutStock(l.ctx, val.ItemId, val.Num)
			if err != nil {
				logx.Errorf("ItemModel.DecutStock: %v, error: %v", val, err)
				return err
			}
			if item.Stock < 0 {
				return dtmcli.ErrFailure
			}
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}

	return &service.DelStockResp{}, nil
}
