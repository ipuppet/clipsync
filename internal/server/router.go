package server

import (
	"clipsync/internal/pkg"
	"encoding/base64"
	"errors"

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
		encoded := base64.StdEncoding.EncodeToString(encrypted)
		handler.JsonStatusWithData(c, encoded, err)
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

		decoded, err := base64.StdEncoding.DecodeString(jsonParam.Data)
		if err != nil {
			handler.JsonStatus(c, errors.New("invalid data"))
			return
		}
		decrypted, err := pkg.AesDecrypt(decoded, []byte(config.Aes.Key), []byte(config.Aes.IV))
		if err != nil {
			handler.JsonStatus(c, err)
			return
		}

		handler.JsonStatus(c, clipboard.WriteAll(string(decrypted)))
	})
}
