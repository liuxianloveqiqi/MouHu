package logic

import (
	"MouHu/service/app/like/rpc/types/like"
	"MouHu/service/common/errorx"
	"context"

	"MouHu/service/app/like/api/internal/svc"
	"MouHu/service/app/like/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeAddLogic {
	return &LikeAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeAddLogic) LikeAdd(req *types.LikeAddReq) error {
	// todo: add your logic here and delete this line
	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return errorx.NewDefaultError("user_id获取失败")
	}
	_, err := l.svcCtx.Rpc.LikeAdd(l.ctx, &like.LikeAddReq{
		LikeType: req.LikeType,
		LikeId:   req.LikeId,
		UserId:   userId,
	})
	if err != nil {
		return errorx.NewDefaultError(err.Error())
	}
	return nil
}
