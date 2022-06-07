package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wazyiz/jwt-gin/controllers"
	"github.com/wazyiz/jwt-gin/middleware"
	"github.com/wazyiz/jwt-gin/models"
)

func main() {
	r := gin.Default()

	models.ConnectDB()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)
	r.Run(":8080")
}
