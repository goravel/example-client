package controllers

import (
	"github.com/goravel/framework/contracts/http"

	servicesuser "goravel/app/services/user"
)

/*********************************
This is the gRPC Client side, if you need the full steps about gRPC, please visit the link below.
https://github.com/goravel/example/blob/master/app/grpc/controllers/user_controller.go
 ********************************/

type UserController struct {
	userService servicesuser.Service
}

func NewUserController() *UserController {
	return &UserController{
		userService: servicesuser.NewServiceImpl(),
	}
}

func (r *UserController) Index(ctx http.Context) {
	token := ctx.Request().Header("Authorization", "")
	if token == "" {
		ctx.Request().AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := r.userService.GetUser(ctx.Context(), token)
	if err != nil {
		ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
			"message": err.Error(),
		})
		return
	}

	ctx.Response().Success().Json(http.Json{
		"user": user,
	})
}
