package auth

import (
	"context"
	"errors"
	"sso/internal/services/auth"
	"sso/internal/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	// Сгенерированный код
	ssov1 "github.com/TaffetaEarth/proto/gen/go/sso"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer // Хитрая штука, о ней ниже
	auth Auth
}

// Тот самый интерфейс, котрый мы передавали в grpcApp
type Auth interface {
	Login(
			ctx context.Context,
			email string,
			password string,
	) (token string, err error)
	RegisterNewUser(
			ctx context.Context,
			email string,
			password string,
	) (userID uint, err error)
}

func Register(gRPCServer *grpc.Server, auth Auth) {  
	ssov1.RegisterAuthServer(gRPCServer, &serverAPI{auth: auth})  
}

func (s *serverAPI) Login(
	ctx context.Context,
	in *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	if in.Email == "" {
			return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if in.Password == "" {
			return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	token, err := s.auth.Login(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
			// Ошибку auth.ErrInvalidCredentials мы создадим ниже
			if errors.Is(err, auth.ErrInvalidCredentials) {
					return nil, status.Error(codes.InvalidArgument, "invalid email or password")
			}

			return nil, status.Error(codes.Internal, "failed to login")
	}

	return &ssov1.LoginResponse{Token: token}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	in *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	if in.Email == "" {
			return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if in.Password == "" {
			return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	uid, err := s.auth.RegisterNewUser(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
			// Ошибку storage.ErrUserExists мы создадим ниже
			if errors.Is(err, storage.ErrUserExists) {
					return nil, status.Error(codes.AlreadyExists, "user already exists")
			}

			return nil, status.Error(codes.Internal, "failed to register user")
	}

	return &ssov1.RegisterResponse{UserId: uint32(uid)}, nil
}
