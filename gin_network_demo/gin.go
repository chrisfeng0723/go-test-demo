/**
 * @Author: fxl
 * @Description:
 * @File:  gin.go
 * @Version: 1.0.0
 * @Date: 2022/8/29 14:26
 */
package gin_network_demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Params struct {
	Name string `json:"name"`
}

func helloHandler(c *gin.Context) {
	var p Params
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "we need a name",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("hello %s", p.Name),
	})

}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/hello", helloHandler)
	return router
}
