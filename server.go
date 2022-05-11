package main

import (
	"zinkworks/grad/server/controller"
	"zinkworks/grad/server/repository"

	log "github.com/inconshreveable/log15"
)

func main() {
	log.Info("Program starting")
	repository.DbInit()
	controller.Run(controller.HandleReqeust())
}
