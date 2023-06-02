package logic

import (
	"MouHu/service/app/qa/model"
	"context"

	"MouHu/service/app/qa/rpc/internal/svc"
	"MouHu/service/app/qa/rpc/types/qa"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelQuestionLogic {
	return &DelQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelQuestionLogic) DelQuestion(in *qa.DelQuestionReq) (*qa.CommonResp, error) {
	// todo: add your logic here and delete this line
	if err := l.svcCtx.Mdb.Delete(&model.Question{}, in.QuestionId).Error; err != nil {
		return nil, err
	}
	return &qa.CommonResp{
		Code:    0,
		Message: "Success!",
	}, nil
}
