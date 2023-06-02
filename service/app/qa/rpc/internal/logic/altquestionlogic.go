package logic

import (
	"MouHu/service/app/qa/model"
	"context"
	"errors"

	"MouHu/service/app/qa/rpc/internal/svc"
	"MouHu/service/app/qa/rpc/types/qa"

	"github.com/zeromicro/go-zero/core/logx"
)

type AltQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAltQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AltQuestionLogic {
	return &AltQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AltQuestionLogic) AltQuestion(in *qa.AltQuestionReq) (*qa.CommonResp, error) {
	// todo: add your logic here and delete this line
	q := model.Question{}
	if result := l.svcCtx.Mdb.Model(&q).Where("id= ?", in.QuestionId).Updates(model.Question{Content: in.Content, Title: in.Title}); result.Error != nil {
		if result.RowsAffected == 0 {
			// 查询结果为空
			return nil, errors.New("查询结果为空")
		} else {
			// 其他错误
			return nil, result.Error
		}
	}
	return &qa.CommonResp{
		Code:    0,
		Message: "Success!",
	}, nil
}
