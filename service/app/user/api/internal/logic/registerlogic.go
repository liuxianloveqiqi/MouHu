package logic

import (
	"MouHu/service/app/user/rpc/types/user"
	"MouHu/service/common/errorx"
	"MouHu/service/common/utils"
	"context"
	"fmt"
	"github.com/google/uuid"

	"MouHu/service/app/user/api/internal/svc"
	"MouHu/service/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.TokenResp, err error) {
	// todo: add your logic here and delete this line
	err = utils.DefaultGetValidParams(l.ctx, req)
	if err != nil {
		return nil, errorx.NewCodeError(100001, fmt.Sprintf("validate校验错误: %v", err))
	}
	cnt, err := l.svcCtx.Rpc.Register(l.ctx, &user.RegisterReq{
		UserPhone: req.UserPhone,
		VeCode:    req.VeCode,
	})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	accessTokenString, refreshTokenString := utils.GetToken(cnt.UserId, uuid.New().String())
	if accessTokenString == "" || refreshTokenString == "" {
		return nil, errorx.NewCodeError(100002, errorx.JWt)
	}

	return &types.TokenResp{
		UserId:       cnt.UserId,
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil

}
