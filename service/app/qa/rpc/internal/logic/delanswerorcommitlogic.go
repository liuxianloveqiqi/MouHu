package logic

import (
	"context"

	"MouHu/service/app/qa/rpc/internal/svc"
	"MouHu/service/app/qa/rpc/types/qa"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAnswerOrCommitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAnswerOrCommitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAnswerOrCommitLogic {
	return &DelAnswerOrCommitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAnswerOrCommitLogic) DelAnswerOrCommit(in *qa.DelAnswerOrCommitReq) (*qa.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &qa.CommonResp{}, nil
}
