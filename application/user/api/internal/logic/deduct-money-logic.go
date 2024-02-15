package logic

import (
	"context"
	"hmall/pkg/util"
	"hmall/pkg/xcode"
	"strconv"

	"hmall/application/user/api/internal/svc"
	"hmall/application/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductMoneyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeductMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductMoneyLogic {
	return &DeductMoneyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeductMoneyLogic) DeductMoney(req *types.DeductMoneyReq) error {
	//1、获得用户Id
	usr, err := util.GetUsr(l.ctx, types.JwtKey)
	if err != nil {
		logx.Errorf("util.GetUsr, error: %v", err)
		return err
	}
	//2、获取用户密码
	user, err := l.svcCtx.UserModel.FindUserById(l.ctx, usr)
	//3、比对密码
	if user.PassWord != util.Md5Password(req.Pw) {
		return xcode.New(types.Unauthorized, "")
	}
	//4、修改金额
	decut, err := strconv.Atoi(req.Amount)
	if err != nil {
		logx.Errorf("strconv.Atoi: %v, error: %v", req.Pw, err)
		return err
	}
	if user.Balance < decut {
		return xcode.New(types.MoneyNotenough, "")
	}
	err = l.svcCtx.UserModel.UpdateBalance(l.ctx, usr, user.Balance-decut)
	if err != nil {
		logx.Errorf("UserModel.UpdateBalance: %v, error: %v", usr, err)
		return err
	}
	//5、返回响应
	return xcode.New(types.OK, "")
}
