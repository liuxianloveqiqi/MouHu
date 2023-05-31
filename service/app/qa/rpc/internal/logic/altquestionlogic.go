package logic

import (
	"context"

	"MouHu/service/app/qa/rpc/internal/svc"
	"MouHu/service/app/qa/rpc/types/qa"

	"github.com/zeromicro/go-zero/core/logx"
)

type AltQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAltQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AltQuestionLogic {
	return &AltQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AltQuestionLogic) AltQuestion(in *qa.AltQuestionReq) (*qa.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &qa.CommonResp{}, nil
}
