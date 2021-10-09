package main

import (
	"code_sim/config"
	"code_sim/es"
	"code_sim/server"
	"flag"
	"os"
)

func main() {
	initConfig()
	es.InitEsCli()
	server.StartServe()
}

func initConfig() {
	var configPath string
	var env string
	flag.StringVar(&configPath, "config_path", "", "配置文件路径")
	flag.StringVar(&env, "env", "", "是否为测试测试环境，值为dev或prd")
	flag.Parse()
	e := func() config.ConfigurationEnv {
		if env == "" {
			var ok bool
			env, ok = os.LookupEnv("ENV")
			if !ok {
				panic("ENV is not set, plz check your environ")
			}
		}
		switch env {
		case "dev":
			return config.DevEnv
		case "prd":
			return config.PrdEnv
		default:
			panic("illegal ENV environ, should be dev or prd")
		}
	}()
	config.InitConfig(configPath, e)
}