package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/threading"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hmall/application/user/rpc/internal/svc"
	"hmall/application/user/rpc/pb"
	"hmall/pkg/util"
	"log"
	"sync"
)

type DecutMoneyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecutMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecutMoneyLogic {
	return &DecutMoneyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecutMoneyLogic) DecutMoney(in *pb.DecutMoneyReq) (*pb.DecutMoneyResp, error) {
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		errorCh := make(chan error, 2)
		wg := &sync.WaitGroup{}
		wg.Add(2)
		//验证密码
		threading.GoSafe(func() {
			defer wg.Done()
			//1、对请求的密码加密
			reqPwd := util.Md5Password(in.Pwd)

			//2、查询密码
			res, err := l.svcCtx.UserModel.FindUserPwdById(l.ctx, in.Uid)
			if err != nil {
				logx.Errorf("UserModel.FindUserPwdById: %v ,error: %v", in.Uid, err)
				errorCh <- dtmcli.ErrFailure
				return
			}
			//3、比对密码
			if res == "" || reqPwd != res {
				log.Println("密码不正确")
				errorCh <- dtmcli.ErrFailure
				return
			}
		})
		//扣减余额
		threading.GoSafe(func() {
			defer wg.Done()
			balance, err := l.svcCtx.UserModel.UpdateBalance(l.ctx, in.Uid, in.Amount)
			if err != nil {
				logx.Errorf("UserModel.UpdateBalance: %v, error: %v", in.Uid, err)
				errorCh <- dtmcli.ErrFailure
				return
			}
			if balance < 0 {
				log.Println("金额不足")
				errorCh <- dtmcli.ErrFailure
				return
			}
		})
		wg.Wait()
		for errCh := range errorCh {
			if errCh != nil {
				return dtmcli.ErrFailure
			}
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}

	return &pb.DecutMoneyResp{}, nil
}
