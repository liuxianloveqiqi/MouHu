package logic

import (
	"MouHu/service/app/like/rpc/types/like"
	"MouHu/service/common/errorx"
	"context"

	"MouHu/service/app/like/api/internal/svc"
	"MouHu/service/app/like/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowUserLogic {
	return &FollowUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowUserLogic) FollowUser(req *types.FollowUserReq) error {
	// todo: add your logic here and delete this line
	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return errorx.NewDefaultError("user_id获取失败")
	}
	_, err := l.svcCtx.Rpc.FollowUser(l.ctx, &like.FollowUserReq{
		FolloweeId: req.FolloweeId,
		UserId:     userId,
	})
	if err != nil {
		return errorx.NewDefaultError(err.Error())
	}
	return nil
}
