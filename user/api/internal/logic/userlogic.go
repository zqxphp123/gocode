package logic

import (
	"context"
	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.UserReq) (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line
	resp = &types.UserResp{
		Data: "ceshis",
	}
	//getuserResp, err := l.svcCtx.User.GetUser(l.ctx, &userclient.UserRequest{
	//	Name: req.Name,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//resp = &types.UserResp{
	//	Data: getuserResp.Data,
	//}
	return resp, nil
}
