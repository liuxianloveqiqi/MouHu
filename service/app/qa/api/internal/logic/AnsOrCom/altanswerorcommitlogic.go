package AnsOrCom

import (
	"context"

	"MouHu/service/app/qa/api/internal/svc"
	"MouHu/service/app/qa/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AltAnswerOrCommitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAltAnswerOrCommitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AltAnswerOrCommitLogic {
	return &AltAnswerOrCommitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AltAnswerOrCommitLogic) AltAnswerOrCommit(req *types.AltAnswerOrCommitReq) error {
	// todo: add your logic here and delete this line

	return nil
}
