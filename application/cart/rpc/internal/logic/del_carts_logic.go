package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/threading"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hmall/application/cart/rpc/internal/svc"
	"hmall/application/cart/rpc/internal/utils"
	"hmall/application/cart/rpc/pb"
	"sync"

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
		wg := &sync.WaitGroup{}
		wg.Add(2)
		threading.GoSafe(func() {
			defer wg.Done()
			if err = l.svcCtx.CartModel.DelCartsByUidItemId(l.ctx, in.Usr, in.ItemId); err != nil {
				logx.Errorf("CartModel.DelCartsByIds: %v, error: %v", in.Usr, err)
			}
		})
		//删缓存
		threading.GoSafe(func() {
			defer wg.Done()
			key := fmt.Sprintf("%s#%d", utils.CacheCartKey, in.Usr)
			_, _ = l.svcCtx.BizRedis.HdelCtx(l.ctx, key, in.ItemId...)
		})
		wg.Wait()

		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}

	return &pb.DelCartsResp{}, nil
}
