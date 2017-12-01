package gotime

import (
	"github.com/bline/gotime/api/proto"
	"github.com/futurenda/google-auth-id-token-verifier"
	"context"
	"strings"
	"errors"
	"time"
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

func NewUserFromIDToken(idtoken string) (*User, error) {
	db := GetDB()
	var curUser User
	set, err := verifyToken(idtoken)
	if err != nil {
		return nil, err
	}
	db.Where("google_id = ?", set.Sub).First(&curUser)
	if curUser.ID != 0 {
		return &curUser, nil
	}
	newUser := &User{
		Email: set.Email,
		GoogleID: set.Sub,
		GivenName: set.GivenName,
		FamilyName: set.FamilyName,
		LastLogin: time.Now(),
		Picture: set.Picture,
		DisplayName: set.Name,
		IsAdmin: false,
	}
	db.Create(&newUser)
	return newUser, nil
}

func verifyToken(token string) (*googleAuthIDTokenVerifier.ClaimSet, error) {
	v := googleAuthIDTokenVerifier.Verifier{}
	aud := "33812767661-4a1p5lotkkveeodjehfpkucvmbpkmkhf.apps.googleusercontent.com"
	err := v.VerifyIDToken(token, []string{
		aud,
	})
	if err != nil {
		return nil, err
	}
	claimSet, err := googleAuthIDTokenVerifier.Decode(token)
	// claimSet.Iss,claimSet.Email ... (See claimset.go)
	// XXX get hd from PrivateClaim?
	if !strings.HasSuffix(claimSet.Email, "@shambhalamountain.org") {
		return nil, errors.New("must use a shambhalamountain.org email address")
	}
	return claimSet, nil
}