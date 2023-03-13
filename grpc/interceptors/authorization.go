package interceptors

import (
	"context"
	"errors"
	"github/achjailani/go-simple-grpc/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

func AuthorizationInterceptor(
	ctx context.Context,
	req interface{},
	_x *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	meta, ok := metadata.FromIncomingContext(ctx)
	if ok == false {
		return nil, errors.New("oh man,.... something happened")
	}

	token, ok := meta["authorization"]

	if !ok {
		return nil, status.Error(codes.Unauthenticated, "No token provided")
	}

	auth := token[0]

	if authorize(auth) == false {
		return nil, status.Error(codes.Unauthenticated, "Invalid username or passwrod")
	}

	return handler(ctx, req)
}

func authorize(token string) bool {
	decoded, err := utils.DecodeBasicAuth(token)

	if err != nil {
		return false
	}

	auth := strings.Split(decoded, ":")
	authUsername := auth[0]
	authPassword := auth[1]

	_ = authUsername
	_ = authPassword

	//us, err := service.NewDBService().User.GetUserByUsername(authUsername)
	//if err != nil {
	//	return false
	//}
	//
	//if us.Password != authPassword {
	//	return false
	//}

	return true
}
