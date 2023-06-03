package logic

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"

	"MouHu/service/app/like/rpc/internal/svc"
	"MouHu/service/app/like/rpc/types/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeAddLogic {
	return &LikeAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LikeAddLogic) LikeAdd(in *like.LikeAddReq) (*like.CommonResp, error) {
	// todo: add your logic here and delete this line
	// 检查用户是否已经点赞
	userId := strconv.FormatInt(in.UserId, 10)
	likeId := strconv.FormatInt(in.LikeId, 10)
	key := "likes:" + in.LikeType + ":" + likeId
	liked, err := l.svcCtx.Rdb.HGet(l.ctx, key, userId).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if liked == "1" {
		return nil, errors.New("已经点赞了无法继续点赞")
	}

	// 在 Hash 中设置用户点赞状态
	err = l.svcCtx.Rdb.HSet(l.ctx, key, userId, "1").Err()
	if err != nil {
		return nil, err
	}

	// 将用户ID添加到已点赞用户集合
	usersKey := key + ":users"
	err = l.svcCtx.Rdb.SAdd(l.ctx, usersKey, userId).Err()
	if err != nil {
		return nil, err
	}
	return &like.CommonResp{
		Code:    0,
		Message: "Success!",
	}, nil
}
