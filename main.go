package main

import (
	"code_sim/config"
	"code_sim/es"
	"code_sim/server"
)

func main() {
	// config.InitConfig()
	config.InitConfigDefault()
	es.InitEsCli()
	server.StartServe()
}
