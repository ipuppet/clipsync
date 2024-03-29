package server

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/gin-gonic/gin"
	"github.com/ipuppet/gtools/handler"
)

func LoadRouters(e *gin.Engine) {
	e.GET("/api/clip", func(c *gin.Context) {
		text, err := clipboard.ReadAll()
		handler.JsonStatusWithData(c, text, err)
	})

	e.POST("/api/clip", func(c *gin.Context) {
		type JsonParam struct {
			Data string `json:"data" binding:"required"`
		}
		var jsonParam JsonParam
		if err := c.BindJSON(&jsonParam); err != nil {
			fmt.Println(jsonParam)
			return
		}

		handler.JsonStatus(c, clipboard.WriteAll(jsonParam.Data))
	})
}
