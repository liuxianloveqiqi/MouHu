package question

import (
	"context"

	"MouHu/service/app/qa/api/internal/svc"
	"MouHu/service/app/qa/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AltQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAltQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AltQuestionLogic {
	return &AltQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AltQuestionLogic) AltQuestion(req *types.AltQuestionReq) error {
	// todo: add your logic here and delete this line

	return nil
}
