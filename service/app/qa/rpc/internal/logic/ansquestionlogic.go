package logic

import (
	"context"

	"MouHu/service/app/qa/rpc/internal/svc"
	"MouHu/service/app/qa/rpc/types/qa"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnsQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnsQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnsQuestionLogic {
	return &AnsQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AnsQuestionLogic) AnsQuestion(in *qa.AnsQuestionReq) (*qa.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &qa.CommonResp{}, nil
}
