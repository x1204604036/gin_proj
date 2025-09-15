package router

import (
	user_router "gin_proj/internal/router/user"

	"github.com/gin-gonic/gin"
)

func LoadRouter(e *gin.Engine) {
	user_router.LoadUserRouter(e)
}
