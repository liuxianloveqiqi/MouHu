package logic

import (
	"context"

	"MouHu/service/app/qa/rpc/internal/svc"
	"MouHu/service/app/qa/rpc/types/qa"

	"github.com/zeromicro/go-zero/core/logx"
)

type AltAnswerOrCommitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAltAnswerOrCommitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AltAnswerOrCommitLogic {
	return &AltAnswerOrCommitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AltAnswerOrCommitLogic) AltAnswerOrCommit(in *qa.AltAnswerOrCommitReq) (*qa.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &qa.CommonResp{}, nil
}
