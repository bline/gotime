package gotime

import (
	"github.com/bline/gotime/api/proto"
	"context"
)

type AccountsService struct {}

func (*AccountsService) GetUser(ctx context.Context, r *api.GetUserRequest) (*api.User, error) {
	return &api.User{}, nil
}
func (*AccountsService) GetUsers(r *api.GetUsersRequest, server api.Accounts_GetUsersServer) error {
	return nil
}
func (*AccountsService) DisableUser(ctx context.Context, r *api.DisableUserRequest) (*api.SimpleResponse, error) {
	return &api.SimpleResponse{IsSuccess: true, Message: "User Disabled"}, nil
}
func (*AccountsService) DeleteUser(ctx context.Context, r *api.DeleteUserRequest) (*api.SimpleResponse, error) {
	return &api.SimpleResponse{IsSuccess: true, Message: "User Deleted"}, nil
}
func (*AccountsService) LockUser(ctx context.Context, r *api.LockUserRequest) (*api.SimpleResponse, error) {
	return &api.SimpleResponse{IsSuccess: true, Message: "User Locked"}, nil
}