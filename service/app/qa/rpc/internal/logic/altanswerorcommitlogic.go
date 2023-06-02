package logic

import (
	"MouHu/service/app/qa/model"
	"context"
	"errors"

	"MouHu/service/app/qa/rpc/internal/svc"
	"MouHu/service/app/qa/rpc/types/qa"

	"github.com/zeromicro/go-zero/core/logx"
)

type AltAnswerOrCommitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAltAnswerOrCommitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AltAnswerOrCommitLogic {
	return &AltAnswerOrCommitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AltAnswerOrCommitLogic) AltAnswerOrCommit(in *qa.AltAnswerOrCommitReq) (*qa.CommonResp, error) {
	// todo: add your logic here and delete this line
	a := model.Answer{}
	if result := l.svcCtx.Mdb.Model(&a).Where("id= ?", in.AnswerOrCommitId).Updates(model.Answer{Content: in.Content}); result.Error != nil {
		if result.RowsAffected == 0 {
			// 查询结果为空
			return nil, errors.New("查询结果为空")
		} else {
			// 其他错误
			return nil, result.Error
		}
	}
	return &qa.CommonResp{}, nil
}
