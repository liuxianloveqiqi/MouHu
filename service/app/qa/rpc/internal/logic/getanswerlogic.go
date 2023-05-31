package logic

import (
	"context"

	"MouHu/service/app/qa/rpc/internal/svc"
	"MouHu/service/app/qa/rpc/types/qa"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnswerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnswerLogic {
	return &GetAnswerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAnswerLogic) GetAnswer(in *qa.GetAnswerReq) (*qa.GetAnswerResp, error) {
	// todo: add your logic here and delete this line

	return &qa.GetAnswerResp{}, nil
}
