package routes

import (
	"clean/controllers"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo, authController *controllers.AuthController) {
	e.POST("/register", authController.Register)
	e.POST("/login", authController.Login)
}
func PostsRoutes(e *echo.Group, postsRoutes *controllers.PostsController) {
	e.GET("", postsRoutes.GetAllPost)
	e.GET("/:id", postsRoutes.GetPostById)
	e.POST("", postsRoutes.CreatePost)
	e.PUT("/:id", postsRoutes.UpdatePost)
	e.DELETE("/:id", postsRoutes.DeletePost)
}
