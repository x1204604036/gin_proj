package main

import (
	"gin_proj/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.LoadRouter(r)
	r.Run(":9898")
}
