package router

import (
	user_controller "gin_proj/internal/controller/user"

	"github.com/gin-gonic/gin"
)

func LoadUserRouter(e *gin.Engine) {
	group := e.Group("/user")
	{
		group.GET("/register", user_controller.UserRegister)
	}
}
