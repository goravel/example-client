package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Web() {
	facades.Route.Get("/", func(ctx http.Context) {
		ctx.Response().Json(200, http.Json{
			"Hello": "Goravel",
		})
	})

	// gRPC
	userController := controllers.NewUserController()
	facades.Route.Get("/user", userController.Index)
}
