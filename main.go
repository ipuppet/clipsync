package main

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
	"github.com/gin-gonic/gin"

	"clipsync/internal/flags"
	"clipsync/internal/icon"
	"clipsync/internal/server"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.GetIconBytes())
	systray.SetTitle("clipsync")
	systray.SetTooltip("clipsync")
	mQuitOrig := systray.AddMenuItem("Quit", "Quit clipsync")
	go func() {
		<-mQuitOrig.ClickedCh
		systray.Quit()
	}()

	server.InitServer(flags.Address)
}

func onExit() {
	fmt.Println("Exit")
	os.Exit(0)
}
