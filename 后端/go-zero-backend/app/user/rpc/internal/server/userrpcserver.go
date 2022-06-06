// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"backend/app/user/rpc/internal/logic"
	"backend/app/user/rpc/internal/svc"
	"backend/app/user/rpc/user"
)

type UserRpcServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserRpcServer
}

func NewUserRpcServer(svcCtx *svc.ServiceContext) *UserRpcServer {
	return &UserRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *UserRpcServer) GetUserInfo(ctx context.Context, in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UserRpcServer) GetUserList(ctx context.Context, in *user.GetUserListReq) (*user.GetUserListResp, error) {
	l := logic.NewGetUserListLogic(ctx, s.svcCtx)
	return l.GetUserList(in)
}
