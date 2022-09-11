package routes

import (
	"jwt-gin/controllers"
	"jwt-gin/middlewares"

	"github.com/gin-gonic/gin"
)

func Handler() {
	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)
	r.Run(":8001")
}
