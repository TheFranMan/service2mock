package main

import (
	"net/http"

	"service2/common"
	"service2/server"

	log "github.com/sirupsen/logrus"
)

func main() {
	env, err := common.GetConfig()
	if nil != err {
		panic(err)
	}

	log.WithField("port", env.Port).Info("Starting server")
	panic(http.ListenAndServe(":"+env.Port, server.New()))
}
