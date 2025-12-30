package main

import (
	"os"

	bootstrap "github.com/leojin/go-service-bootstrap"

	serverAdmin "github.com/leojin/go-service-sample/server/admin"
	serverBackend "github.com/leojin/go-service-sample/server/backend"
)

func main() {

	cwd, _ := os.Getwd()

	boot := bootstrap.ServiceBootstrap{}
	if err := boot.Init(cwd); err != nil {
		os.Exit(1)
	}

	boot.RegisterServer(serverAdmin.NewServer())
	boot.RegisterServer(serverBackend.NewServer())

	boot.Start()
}
