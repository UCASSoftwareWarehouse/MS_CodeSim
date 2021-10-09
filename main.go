package main

import (
	"code_sim/config"
	"code_sim/es"
	"code_sim/server"
)

func main() {
	config.InitConfig()
	es.InitEsCli()
	server.StartServe()
}
