package commit

import (
	"context"

	"MouHu/service/app/qa/api/internal/svc"
	"MouHu/service/app/qa/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ComAnswerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewComAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ComAnswerLogic {
	return &ComAnswerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ComAnswerLogic) ComAnswer(req *types.ComAnswerReq) error {
	// todo: add your logic here and delete this line

	return nil
}
