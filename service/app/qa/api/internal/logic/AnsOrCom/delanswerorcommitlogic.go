package AnsOrCom

import (
	"context"

	"MouHu/service/app/qa/api/internal/svc"
	"MouHu/service/app/qa/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAnswerOrCommitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelAnswerOrCommitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAnswerOrCommitLogic {
	return &DelAnswerOrCommitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelAnswerOrCommitLogic) DelAnswerOrCommit(req *types.DelAnswerOrCommitReq) error {
	// todo: add your logic here and delete this line

	return nil
}
