package controller

import (
	"encoding/json"
	"fmt"
	"gin_proj/pkg/common_struct"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	registerInfo := &common_struct.RegisterInfo{}
	err := c.ShouldBind(registerInfo)
	if err != nil {
		c.JSON(400, gin.H{"code": -1, "msg": "param bind error"})
		return
	}

	_registerInfo, err := json.Marshal(registerInfo)
	fmt.Println("registerInfo: ", string(_registerInfo), ", error: ", err)

	c.JSON(200, gin.H{"code": 0, "msg": "success"})
}
