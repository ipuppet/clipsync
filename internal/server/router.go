package server

import (
	"clipsync/internal/pkg"

	"github.com/atotto/clipboard"
	"github.com/gin-gonic/gin"
	"github.com/ipuppet/gtools/handler"
)

func LoadRouters(e *gin.Engine) {
	e.HEAD("/api/ping", func(c *gin.Context) {
		handler.JsonStatusWithData(c, "pong", nil)
	})

	e.GET("/api/clip", func(c *gin.Context) {
		text, err := clipboard.ReadAll()
		if err != nil {
			handler.JsonStatus(c, err)
			return
		}

		encrypted, err := pkg.AesEncrypt([]byte(text), []byte(config.Aes.Key), []byte(config.Aes.IV))
		handler.JsonStatusWithData(c, string(encrypted), err)
	})

	e.POST("/api/clip", func(c *gin.Context) {
		type JsonParam struct {
			Data string `json:"data" binding:"-"`
		}
		var jsonParam JsonParam
		if err := c.BindJSON(&jsonParam); err != nil {
			handler.JsonStatus(c, err)
			return
		}

		decrypted, err := pkg.AesDecrypt([]byte(jsonParam.Data), []byte(config.Aes.Key), []byte(config.Aes.IV))
		if err != nil {
			handler.JsonStatus(c, err)
			return
		}

		handler.JsonStatus(c, clipboard.WriteAll(string(decrypted)))
	})
}
