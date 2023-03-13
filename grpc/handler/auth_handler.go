package handler

import (
	"context"
	"github/achjailani/go-simple-grpc/proto/foo"
	"github/achjailani/go-simple-grpc/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthLogin is a function
func (c *Handler) AuthLogin(ctx context.Context, r *foo.AuthLoginPayload) (*foo.LoginResponse, error) {
	username := r.GetUsername()
	password := r.GetPassword()

	us, err := c.repo.User.FindByUsername(ctx, username)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid username or password")
	}

	if us.Password != password {
		return nil, status.Error(codes.InvalidArgument, "Invalid username or password")
	}

	return &foo.LoginResponse{
		Ok:          true,
		AccessToken: utils.EncodeBasicAuth(us.Username, us.Password),
	}, nil
}
