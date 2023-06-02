package answer

import (
	"MouHu/service/app/qa/rpc/types/qa"
	"MouHu/service/common/errorx"
	"context"

	"MouHu/service/app/qa/api/internal/svc"
	"MouHu/service/app/qa/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnswerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnswerLogic {
	return &GetAnswerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAnswerLogic) GetAnswer() (resp *types.GetAnswerResp, err error) {
	// todo: add your logic here and delete this line
	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return nil, errorx.NewDefaultError("user_id获取失败")
	}
	res, err := l.svcCtx.Rpc.GetAnswer(l.ctx, &qa.GetAnswerReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	alist := make([]types.AnswerList, 0)
	for _, v := range res.Answer {
		a := types.AnswerList{
			AnswerId:   v.AnswerId,
			UserId:     v.UserId,
			QuestionId: v.QuestionId,
			ParentId:   v.ParentId,
			Content:    v.Content,
			IsAnswer:   v.IsAnswer,
			CreateTime: v.CreateTime.AsTime().Format("2006-01-02 15:04:05"),
			UpdateTime: v.UpdateTime.AsTime().Format("2006-01-02 15:04:05"),
			DeleteTime: v.DeleteTime.AsTime().Format("2006-01-02 15:04:05"),
		}
		alist = append(alist, a)
	}
	return &types.GetAnswerResp{Answer: alist}, nil
}
