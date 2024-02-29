package logic

import (
	"context"
	"database/sql"
	"fmt"
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
		chErr := make(chan error, 2)
		wg := &sync.WaitGroup{}
		wg.Add(2)
		threading.GoSafe(func() {
			defer wg.Done()
			if err = l.svcCtx.CartModel.DelCartsByUidItemId(l.ctx, in.Usr, in.ItemId); err != nil {
				logx.Errorf("CartModel.DelCartsByIds: %v, error: %v", in.Usr, err)
				chErr <- status.Error(codes.Internal, err.Error())
			}
		})
		//删缓存
		threading.GoSafe(func() {
			defer wg.Done()
			key := fmt.Sprintf("%s#%d", utils.CacheCartKey, in.Usr)
			if _, err = l.svcCtx.BizRedis.Hdel(key); err != nil {
				logx.Errorf("BizRedis.Hdel: %v, error: %v", key, err)
				chErr <- status.Error(codes.Internal, err.Error())
			}
		})
		wg.Wait()
		close(chErr)

		for e := range chErr {
			if e == status.Error(codes.Internal, err.Error()) {
				return e
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.DelCartsResp{}, nil
}
