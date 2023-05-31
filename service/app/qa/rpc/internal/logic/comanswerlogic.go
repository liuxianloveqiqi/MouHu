package logic

import (
	"context"

	"MouHu/service/app/qa/rpc/internal/svc"
	"MouHu/service/app/qa/rpc/types/qa"

	"github.com/zeromicro/go-zero/core/logx"
)

type ComAnswerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewComAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ComAnswerLogic {
	return &ComAnswerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ComAnswerLogic) ComAnswer(in *qa.ComAnswerReq) (*qa.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &qa.CommonResp{}, nil
}
