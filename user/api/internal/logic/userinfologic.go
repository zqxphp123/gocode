package logic

import (
	"context"
	"user/rpc/userclient"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserinfoLogic) Userinfo(req *types.UserReq) (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line
	getuserResp, err := l.svcCtx.User.GetUser(l.ctx, &userclient.UserRequest{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.UserResp{
		Data: getuserResp.Data,
	}
	return resp, nil

}
