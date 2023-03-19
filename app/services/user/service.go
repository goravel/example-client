package user

import (
	"context"

	proto "github.com/goravel/example-proto"
)

type Service interface {
	GetUser(ctx context.Context, token string) (*proto.User, error)
}
