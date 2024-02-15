package logic

import (
	"context"
	"hmall/application/user/api/internal/model"
	"hmall/pkg/jwt"
	"hmall/pkg/util"
	"hmall/pkg/xcode"
	"log"

	"hmall/application/user/api/internal/svc"
	"hmall/application/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	//1、对请求的密码加密
	reqPwd := util.Md5Password(req.Password)

	//2、查询密码
	var res model.UserDTO
	res, err = l.svcCtx.UserModel.FindUserByName(l.ctx, req.UserName)
	if err != nil {
		log.Println("UserModel.FindUserByName: %v ,error: %v", req.UserName, err)
		return nil, err
	}
	if res.Id == 0 {
		return nil, xcode.New(types.NotFound, "")
	}
	//3、比对密码
	if reqPwd != res.PassWord {
		return nil, xcode.New(types.Unauthorized, "")
	}
	//4、生成jwt令牌
	opt := &jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			types.JwtKey: res.Id,
		},
	}
	token, err := jwt.BuildTokens(*opt)
	if err != nil {
		logx.Errorf("BuildTokens, error: %v", err)
		return nil, err
	}
	//5、返回响应
	return &types.LoginResp{
		Token:    token,
		Balance:  res.Balance,
		UserName: res.UserName,
		UserId:   res.Id,
	}, nil
}
