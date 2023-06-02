package logic

import (
	"MouHu/service/app/qa/model"
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	"MouHu/service/app/qa/rpc/internal/svc"
	"MouHu/service/app/qa/rpc/types/qa"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnswerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnswerLogic {
	return &GetAnswerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAnswerLogic) GetAnswer(in *qa.GetAnswerReq) (*qa.GetAnswerResp, error) {
	// todo: add your logic here and delete this line
	var answerList []model.Answer
	if result := l.svcCtx.Mdb.Where("user_id = ? and is_answer = ?", in.UserId, 1).Find(&answerList); result.Error != nil {
		if result.RowsAffected == 0 {
			// 查询结果为空
			return nil, errors.New("查询结果为空")
		} else {
			// 其他错误
			return nil, result.Error
		}
	}
	alist := make([]*qa.AnswerList, 0)
	for _, v := range answerList {
		a := &qa.AnswerList{
			AnswerId:   v.Id,
			UserId:     v.UserId,
			QuestionId: v.QuestionId,
			ParentId:   v.ParentId,
			Content:    v.Content,
			IsAnswer:   true,
			CreateTime: timestamppb.New(v.CreateTime),
			UpdateTime: timestamppb.New(v.UpdateTime),
			DeleteTime: timestamppb.New(v.DeleteTime.Time),
		}
		alist = append(alist, a)
	}
	return &qa.GetAnswerResp{Answer: alist}, nil
}
