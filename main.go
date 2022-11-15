package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/atotto/clipboard"
	"github.com/gin-gonic/gin"
	"github.com/ipuppet/gtools/handler"
	"golang.org/x/sync/errgroup"
)

func GetServer(addr string, handle func(engine *gin.Engine)) *http.Server {
	engine := gin.Default()

	engine.Use(func(c *gin.Context) {
		c.Next()

		// check error
		for _, err := range c.Errors {
			c.JSON(c.Writer.Status(), gin.H{
				"status":  false,
				"message": err.Error(),
			})
			c.Abort()

			return
		}
	})

	handle(engine)

	fmt.Println("server listening on: " + addr)

	return &http.Server{
		Addr:         addr,
		Handler:      engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	host := os.Getenv("GO_HOST")

	server := GetServer(host+":8080", func(engine *gin.Engine) {
		LoadRouters(engine)
	})

	var g errgroup.Group
	g.Go(func() error {
		return server.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func LoadRouters(e *gin.Engine) {
	e.GET("/api/clip", func(c *gin.Context) {
		text, err := clipboard.ReadAll()
		handler.JsonStatusWithData(c, text, err)
	})

	e.POST("/api/clip", func(c *gin.Context) {
		type JsonParam struct {
			Content string `form:"content" json:"content" binding:"required"`
		}
		var jsonParam JsonParam
		if err := c.ShouldBind(&jsonParam); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		handler.JsonStatus(c, clipboard.WriteAll(jsonParam.Content))
	})
}
