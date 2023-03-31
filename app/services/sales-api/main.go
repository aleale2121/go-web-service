package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	_ "go.uber.org/automaxprocs"
	"go.uber.org/automaxprocs/maxprocs"
)

var build = "develop"

func main() {
	//	github.com/ardanlabs/conf v1.5.0
	// github.com/dimfeld/httptreemux/v5 v5.5.0
	if _, err := maxprocs.Set(); err != nil {
		fmt.Println("maxprocs: %w", err)
		os.Exit(1)
	}
	g := runtime.GOMAXPROCS(0)
	log.Printf("Starting service build[%s] CPU[%d]", build, g)
	defer log.Println("service ended")
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	log.Println("Stopping service")
}
