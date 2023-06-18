package user

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	proto "github.com/goravel/example-proto"
	"github.com/goravel/framework/facades"
)

var instance *ServiceImpl
var once sync.Once

type ServiceImpl struct {
	userServiceClient proto.UserServiceClient
	timeout           time.Duration
}

func NewServiceImpl() *ServiceImpl {
	once.Do(func() {
		client, err := facades.Grpc().Client(context.Background(), "user")
		if err != nil {
			facades.Log().Errorf("init UserService err: %+v", err)
			return
		}

		instance = &ServiceImpl{
			userServiceClient: proto.NewUserServiceClient(client),
			timeout:           10 * time.Second,
		}
	})

	return instance
}

func (r *ServiceImpl) GetUser(ctx context.Context, token string) (*proto.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := r.userServiceClient.GetUser(ctx, &proto.UserRequest{
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != http.StatusOK {
		return nil, fmt.Errorf("user service returns error, code: %d, message: %s", res.Code, res.Message)
	}

	return res.GetData(), nil
}
