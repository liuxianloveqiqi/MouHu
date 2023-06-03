package logic

import (
	"MouHu/service/app/like/model"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"

	"MouHu/service/app/like/rpc/internal/svc"
	"MouHu/service/app/like/rpc/types/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowUserLogic {
	return &FollowUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowUserLogic) FollowUser(in *like.FollowUserReq) (*like.CommonResp, error) {
	// todo: add your logic here and delete this line
	existingFollow := model.Follow{}
	err := l.svcCtx.Mdb.Where("follower_id = ? and followee_id = ?", in.UserId, in.FolloweeId).First(&existingFollow).Error
	if err == nil {
		// 已存在关注记录，返回错误
		return nil, errors.New("已关注该用户")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 查询过程中发生错误，返回错误
		logx.Errorf("查询关注记录失败: %v", err)
		return nil, err
	}

	f := model.Follow{
		FollowerId: in.UserId,
		FolloweeId: in.FolloweeId,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := l.svcCtx.Mdb.Create(&f).Error; err != nil {
		logx.Errorf("创建记录失败: %v", err)
		return nil, err
	}

	return &like.CommonResp{
		Code:    0,
		Message: "Success!",
	}, nil
}
