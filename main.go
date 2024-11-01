package main

import (
	"clean/config"
	"clean/controllers"
	"clean/repository"
	"clean/routes"
	"clean/services"
	"log"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	db, _ := config.InitDB()

	// ini autentikasi
	userRepo := repository.NewUserRepo(db)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	// ini posts
	postRepo := repository.NewPostsRepository(db)
	postService := services.NewPostsService(postRepo)
	postController := controllers.NewPostsController(postService)

	e := echo.New()

	// rutes auth
	routes.AuthRoutes(e, authController)

	// seting jwt dan gunakan routes group
	eauth := e.Group("/posts")
	eauth.Use(echojwt.JWT([]byte("mono")))
	routes.PostsRoutes(eauth, postController)
	e.Start(":8000")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("failed lod env")
	}
}
