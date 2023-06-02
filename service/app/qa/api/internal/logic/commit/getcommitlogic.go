package commit

import (
	"context"

	"MouHu/service/app/qa/api/internal/svc"
	"MouHu/service/app/qa/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommitLogic {
	return &GetCommitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommitLogic) GetCommit() (resp *types.GetCommitResp, err error) {
	// todo: add your logic here and delete this line

	return
}
