// Code generated by goctl. DO NOT EDIT.
// Source: like.proto

package likeclient

import (
	"context"

	"MouHu/service/app/like/rpc/types/like"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CommonResp    = like.CommonResp
	FollowUserReq = like.FollowUserReq
	LikeAddReq    = like.LikeAddReq

	Like interface {
		FollowUser(ctx context.Context, in *FollowUserReq, opts ...grpc.CallOption) (*CommonResp, error)
		LikeAdd(ctx context.Context, in *LikeAddReq, opts ...grpc.CallOption) (*CommonResp, error)
	}

	defaultLike struct {
		cli zrpc.Client
	}
)

func NewLike(cli zrpc.Client) Like {
	return &defaultLike{
		cli: cli,
	}
}

func (m *defaultLike) FollowUser(ctx context.Context, in *FollowUserReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := like.NewLikeClient(m.cli.Conn())
	return client.FollowUser(ctx, in, opts...)
}

func (m *defaultLike) LikeAdd(ctx context.Context, in *LikeAddReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := like.NewLikeClient(m.cli.Conn())
	return client.LikeAdd(ctx, in, opts...)
}
