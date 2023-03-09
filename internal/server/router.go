package server

import (
	"net/http"

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
			Data string `form:"data" json:"data" binding:"required"`
		}
		var jsonParam JsonParam
		if err := c.ShouldBind(&jsonParam); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		handler.JsonStatus(c, clipboard.WriteAll(jsonParam.Data))
	})
}
