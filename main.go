package main

import (
	"fmt"
	"os"
	"os/signal"

	"dtech.com/traktirku/cmd"
)

func main() {
	cmd.RunApp()
	fmt.Println("Server starts in background...")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
