package auth

import (
	"context"

	ssov1 "github.com/renlin-code/grpc-sso-microservice_protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	SignIn(ctx context.Context, email, password string, appId int) (string, error)
	CreateNewUser(ctx context.Context, email, password string) (int, error)
	IsAdmin(ctx context.Context, userId int) (bool, error)
}

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

func (s *serverAPI) SignUp(ctx context.Context, req *ssov1.SignUpRequest) (*ssov1.SignUpResponse, error) {
	if err := validateSignUp(req); err != nil {
		return nil, err
	}

	userId, err := s.auth.CreateNewUser(ctx, req.GetEmail(), req.GetPassword())

	if err != nil {
		//TODO: handle error
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &ssov1.SignUpResponse{
		UserId: int64(userId),
	}, nil
}

func (s *serverAPI) SignIn(ctx context.Context, req *ssov1.SignInRequest) (*ssov1.SignInResponse, error) {
	if err := validateSignIn(req); err != nil {
		return nil, err
	}

	token, err := s.auth.SignIn(ctx, req.GetEmail(), req.GetPassword(), int(req.GetAppId()))

	if err != nil {
		//TODO: handle error
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &ssov1.SignInResponse{Token: token}, nil
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	if err := validateIsAdmin(req); err != nil {
		return nil, err
	}

	isAdmin, err := s.auth.IsAdmin(ctx, int(req.GetUserId()))
	if err != nil {
		//TODO: handle error
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssov1.IsAdminResponse{IsAdmin: isAdmin}, nil
}

// TODO: make a better validation
func validateSignUp(req *ssov1.SignUpRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}
	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	return nil
}

func validateSignIn(req *ssov1.SignInRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}
	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	if req.GetAppId() == 0 {
		return status.Error(codes.InvalidArgument, "app_id is required")
	}
	return nil
}

func validateIsAdmin(req *ssov1.IsAdminRequest) error {
	if req.UserId == 0 {
		return status.Error(codes.InvalidArgument, "user_id is required")
	}
	return nil
}
