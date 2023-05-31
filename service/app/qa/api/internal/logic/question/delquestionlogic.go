package question

import (
	"context"

	"MouHu/service/app/qa/api/internal/svc"
	"MouHu/service/app/qa/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelQuestionLogic {
	return &DelQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelQuestionLogic) DelQuestion(req *types.DelQuestionReq) error {
	// todo: add your logic here and delete this line

	return nil
}
