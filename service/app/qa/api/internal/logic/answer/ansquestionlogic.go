package answer

import (
	"MouHu/service/app/qa/rpc/types/qa"
	"MouHu/service/common/errorx"
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
	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return errorx.NewDefaultError("user_id获取失败")
	}

	_, err := l.svcCtx.Rpc.AnsQuestion(l.ctx, &qa.AnsQuestionReq{
		QuestionId: req.QuestionId,
		Content:    req.Content,
		UserId:     userId,
	})
	if err != nil {

		return errorx.NewDefaultError(err.Error())
	}
	return nil
}
