package logic

import (
	"MouHu/service/app/qa/model"
	"MouHu/service/app/qa/rpc/internal/svc"
	"MouHu/service/app/qa/rpc/types/qa"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionLogic {
	return &GetQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetQuestionLogic) GetQuestion(in *qa.GetQuestionReq) (*qa.GetQuestionResp, error) {
	// todo: add your logic here and delete this line
	var questionList []model.Question

	if result := l.svcCtx.Mdb.Where("user_id = ?", in.UserId).Find(&questionList); result.Error != nil {
		if result.RowsAffected == 0 {
			// 查询结果为空
			return nil, errors.New("查询结果为空")
		} else {
			// 其他错误
			return nil, result.Error
		}
	}
	qlist := make([]*qa.QuestionList, 0)
	for _, v := range questionList {
		q := &qa.QuestionList{
			QuestionId: v.Id,
			UserId:     v.UserId,
			Title:      v.Title,
			Content:    v.Content,
			CreateTime: timestamppb.New(v.CreateTime),
			UpdateTime: timestamppb.New(v.UpdateTime),
			DeleteTime: timestamppb.New(v.DeleteTime.Time),
		}

		qlist = append(qlist, q)
	}

	return &qa.GetQuestionResp{Question: qlist}, nil
}
