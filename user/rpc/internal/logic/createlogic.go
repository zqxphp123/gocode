package logic

import (
	"context"
	"database/sql"
	"time"
	"user/model"

	"user/rpc/internal/svc"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *user.CreateRep) (*user.CreateResp, error) {
	// todo: add your logic here and delete this line
	var name sql.NullString
	if in.Name != "" {
		name.String = in.Name
		name.Valid = true
	}
	_, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Id:       int64(in.Id),
		Name:     name,
		Password: in.Phone,
		Mobile:   in.Phone,
		Gender:   "",
		Nickname: "ceshi",
		Type:     0,
		CreateAt: sql.NullTime{},
		UpdateAt: time.Time{},
	})
	return &user.CreateResp{}, err
}
