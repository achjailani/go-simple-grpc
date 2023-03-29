package handler

import (
	"context"
	"github/achjailani/go-simple-grpc/domain/entity"
	"github/achjailani/go-simple-grpc/proto/foo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetUser is a function
func (c *Handler) GetUser(ctx context.Context, r *foo.UserByIDRequest) (*foo.User, error) {
	usr, err := c.dep.repo.User.Find(ctx, int(r.GetId()))

	if err != nil {
		return nil, status.Error(codes.NotFound, "Data Not Found")
	}

	return &foo.User{
		Id:        uint64(usr.ID),
		Name:      usr.Name,
		Username:  usr.Username,
		CreatedAt: usr.CreatedAt.String(),
		UpdatedAt: usr.UpdatedAt.String(),
		DeletedAt: usr.DeletedAt.Time.String(),
	}, nil
}

// UpdateUser is function
func (c *Handler) UpdateUser(ctx context.Context, payload *foo.UserUpdateRequest) (*foo.User, error) {
	userId := int(payload.GetId())

	if _, err := c.dep.repo.User.Find(ctx, userId); err != nil {
		return nil, status.Error(codes.NotFound, "Data Not Found")
	}

	userData := &entity.User{
		Name:     payload.GetName(),
		Username: payload.GetUsername(),
	}

	err := c.dep.repo.User.Update(ctx, userId, userData)
	if err != nil {
		return nil, err
	}

	return &foo.User{
		Id:        uint64(userData.ID),
		Name:      userData.Name,
		Username:  userData.Username,
		CreatedAt: userData.CreatedAt.String(),
		UpdatedAt: userData.UpdatedAt.String(),
		DeletedAt: userData.CreatedAt.String(),
	}, nil
}

// CreateUser is function
func (c *Handler) CreateUser(ctx context.Context, r *foo.UserCreateRequest) (*foo.User, error) {
	usr := entity.User{
		Name:     r.GetName(),
		Username: r.GetUsername(),
		Password: r.GetPassword(),
	}

	err := c.dep.repo.User.Create(ctx, &usr)

	if err != nil {
		return nil, err
	}

	return &foo.User{
		Id:        uint64(usr.ID),
		Name:      usr.Name,
		Username:  usr.Username,
		CreatedAt: usr.CreatedAt.String(),
		UpdatedAt: usr.UpdatedAt.String(),
		DeletedAt: usr.CreatedAt.String(),
	}, nil
}

// GetUserList is function
func (c *Handler) GetUserList(ctx context.Context, _ *foo.UserListQuery) (*foo.Users, error) {
	serv, err := c.dep.repo.User.Get(ctx)

	if err != nil {
		return nil, err
	}

	var users []*foo.User
	for _, u := range serv {
		users = append(users, &foo.User{
			Id:        uint64(u.ID),
			Name:      u.Name,
			Username:  u.Username,
			CreatedAt: u.CreatedAt.String(),
			UpdatedAt: u.UpdatedAt.String(),
			DeletedAt: u.CreatedAt.String(),
		})
	}

	return &foo.Users{
		Users: users,
	}, nil
}

// DeleteUser is a function
func (c *Handler) DeleteUser(ctx context.Context, r *foo.UserByIDRequest) (*foo.UserDeleteResponse, error) {
	userId := int(r.GetId())

	if _, err := c.dep.repo.User.Find(ctx, userId); err != nil {
		return nil, status.Error(codes.NotFound, "Data not found")
	}

	err := c.dep.repo.User.Delete(ctx, int(r.GetId()))

	if err != nil {
		return nil, err
	}

	return &foo.UserDeleteResponse{
		Message: "ok",
	}, nil
}
