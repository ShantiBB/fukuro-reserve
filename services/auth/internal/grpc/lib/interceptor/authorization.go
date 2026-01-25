package interceptor

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	userv1 "auth/api/user/v1"
	"auth/internal/grpc/lib/utils/helper"
	"auth/internal/grpc/lib/utils/permission"
	"auth/pkg/lib/utils/consts"
	"auth/pkg/lib/utils/jwt"
)

type contextKey string

const (
	ClaimsKey          contextKey = "claims"
	methodRegisterUser string     = "/user.v1.TokenService/RegisterUser"
	methodLoginUser    string     = "/user.v1.TokenService/LoginUser"
	methodRefreshToken string     = "/user.v1.TokenService/RefreshToken"
)

func AuthInterceptor(accessSecret string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		publicMethods := map[string]struct{}{
			methodRegisterUser: {},
			methodLoginUser:    {},
			methodRefreshToken: {},
		}

		if _, ok := publicMethods[info.FullMethod]; ok {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, helper.HandleDomainErr(consts.ErrUnauthorized)
		}

		token := md.Get("authorization")
		if len(token) == 0 {
			return nil, helper.HandleDomainErr(consts.ErrUnauthorized)
		}

		bearerToken, ok := strings.CutPrefix(token[0], "Bearer ")
		if !ok {
			return nil, status.Error(codes.Unauthenticated, consts.MsgInvalidBearer)
		}

		claims, err := jwt.ParseBearerToken(bearerToken, accessSecret)
		if err != nil {
			return nil, helper.HandleDomainErr(consts.ErrInvalidCredentials)
		}

		userID := extractUserID(req)
		if ok = permission.CheckRolePermission(info.FullMethod, claims.Role, claims.Sub, userID); !ok {
			return nil, helper.HandleDomainErr(consts.ErrForbidden)
		}

		ctx = context.WithValue(ctx, ClaimsKey, claims)

		return handler(ctx, req)
	}
}

func extractUserID(req any) int64 {
	switch r := req.(type) {
	case *userv1.GetUserRequest:
		return r.Id
	case *userv1.UpdateUserRequest:
		return r.Id
	case *userv1.DeleteUserRequest:
		return r.Id
	default:
		return 0
	}
}
