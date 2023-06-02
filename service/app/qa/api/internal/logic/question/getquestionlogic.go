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

func (l *GetQuestionLogic) GetQuestion() (resp *types.GetQuestionResp, err error) {
	// todo: add your logic here and delete this line
	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return nil, errorx.NewDefaultError("user_id获取失败")
	}
	res, err := l.svcCtx.Rpc.GetQuestion(l.ctx, &qa.GetQuestionReq{
		UserId: userId,
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
			CreateTime: v.CreateTime.AsTime().Format("2006-01-02 15:04:05"),
			UpdateTime: v.UpdateTime.AsTime().Format("2006-01-02 15:04:05"),
			DeleteTime: v.DeleteTime.AsTime().Format("2006-01-02 15:04:05"),
		}
		qlist = append(qlist, q)
	}
	return &types.GetQuestionResp{Question: qlist}, nil
}
