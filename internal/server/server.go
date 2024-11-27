package server

import (
	"clipsync/internal/flags"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func getServer(addr string) *http.Server {
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

	LoadRouters(engine)

	fmt.Println("server listening on " + addr)

	return &http.Server{
		Addr:         addr,
		Handler:      engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

}

func InitServer() {
	server := getServer(flags.Address + ":" + flags.Port)

	var g errgroup.Group
	g.Go(func() error {
		return server.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
