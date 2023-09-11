package controllers

import (
	"github.com/goravel/framework/contracts/http"

	servicesuser "goravel/app/services/user"
)

/*********************************
gRPC Example

This is the gRPC Client side, if you need the full steps about gRPC, please visit the link below.
https://github.com/goravel/example/blob/master/app/grpc/controllers/user_controller.go

There is a client interceptor example about opentracing, you can find it in `app/grpc/interceptors/opentracing_client.go`
[gRPC Document](https://www.goravel.dev/the-basics/grpc.html)
 ********************************/

type UserController struct {
	userService servicesuser.Service
}

func NewUserController() *UserController {
	return &UserController{
		userService: servicesuser.NewServiceImpl(),
	}
}

func (r *UserController) Index(ctx http.Context) http.Response {
	token := ctx.Request().Header("Authorization", "")
	if token == "" {
		return ctx.Response().Status(http.StatusUnauthorized).String("")
	}

	user, err := r.userService.GetUser(ctx.Context(), token)
	if err != nil {
		return ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"message": err.Error(),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"user": user,
	})
}
