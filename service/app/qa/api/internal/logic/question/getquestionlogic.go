package question

import (
	"MouHu/service/app/qa/rpc/types/qa"
	"MouHu/service/common/errorx"
	"context"

	"MouHu/service/app/qa/api/internal/svc"
	"MouHu/service/app/qa/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionLogic {
	return &GetQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetQuestionLogic) GetQuestion(req *types.GetQuestionReq) (resp *types.GetQuestionResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.Rpc.GetQuestion(l.ctx, &qa.GetQuestionReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	qlist := make([]types.QuestionList, 0)
	for _, v := range res.Question {
		q := types.QuestionList{
			QuestionId: v.QuestionId,
			UserID:     v.UserId,
			Title:      v.Title,
			Content:    v.Content,
			CreateTime: v.CreateTime.String(),
			UpdateTime: v.UpdateTime.String(),
			DeleteTime: v.DeleteTime.String(),
		}
		qlist = append(qlist, q)
	}
	return &types.GetQuestionResp{Question: qlist}, nil
}
