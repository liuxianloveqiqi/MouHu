package question

import (
	"MouHu/service/app/qa/rpc/types/qa"
	"MouHu/service/common/errorx"
	"MouHu/service/common/utils"
	"context"
	"fmt"

	"MouHu/service/app/qa/api/internal/svc"
	"MouHu/service/app/qa/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PubQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPubQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PubQuestionLogic {
	return &PubQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PubQuestionLogic) PubQuestion(req *types.PubQuestionReq) error {
	// todo: add your logic here and delete this line
	err := utils.DefaultGetValidParams(l.ctx, req)
	if err != nil {
		return errorx.NewDefaultError(fmt.Sprintf("validate校验错误: %v", err))
	}
	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return errorx.NewDefaultError("user_id获取失败")
	}

	_, err = l.svcCtx.Rpc.PubQuestion(l.ctx, &qa.PubQuestionReq{
		Title:   req.Title,
		Content: req.Content,
		UserId:  userId,
	})
	if err != nil {

		return errorx.NewDefaultError(err.Error())
	}
	return nil
}
