package answer

import (
	"context"

	"MouHu/service/app/qa/api/internal/svc"
	"MouHu/service/app/qa/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnsQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAnsQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnsQuestionLogic {
	return &AnsQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AnsQuestionLogic) AnsQuestion(req *types.AnsQuestionReq) error {
	// todo: add your logic here and delete this line

	return nil
}
