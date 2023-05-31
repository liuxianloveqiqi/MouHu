package logic

import (
	"context"

	"MouHu/service/app/qa/rpc/internal/svc"
	"MouHu/service/app/qa/rpc/types/qa"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommitLogic {
	return &GetCommitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommitLogic) GetCommit(in *qa.GetCommitReq) (*qa.GetCommitResp, error) {
	// todo: add your logic here and delete this line

	return &qa.GetCommitResp{}, nil
}
