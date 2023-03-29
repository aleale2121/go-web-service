package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ardanlabs/conf"
	"github.com/dimfeld/httptreemux/v5"
)

var build = "develop"

func main() {
	_=conf.ErrHelpWanted
	log.Println("Starting service ", build)
	defer log.Println("service ended")
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	log.Println("Stopping service")
	httptreemux.New()
}
