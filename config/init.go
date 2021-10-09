package config

import (
	"flag"
	"os"
)

var Conf *EachConfig
var cmdArgs *CmdArgs
var environArgs *EnvironArgs

func InitConfig() {
	initCmdArgs()
	initEnvironArgs()
	configs := parse()
	Conf = configs[cmdArgs.Env]
	Conf.Env = cmdArgs.Env
	setConfByEnvirons()
}

func initCmdArgs() {
	var configPath string
	var env string
	var port int
	flag.StringVar(&configPath, "config_path", "", "配置文件路径")
	flag.StringVar(&env, "env", "", "是否为测试测试环境，值为dev或prd")
	flag.IntVar(&port, "port", 0, "指定端口号，默认为配置文件中配置的端口号")
	flag.Parse()
	e := func() ConfigurationEnv {
		if env == "" {
			var ok bool
			env, ok = os.LookupEnv("ENV")
			if !ok {
				panic("ENV is not set, plz check your environ")
			}
		}
		switch env {
		case "dev":
			return DevEnv
		case "prd":
			return PrdEnv
		default:
			panic("illegal ENV environ, should be dev or prd")
		}
	}()
	cmdArgs = &CmdArgs{
		Env:        e,
		Port:       port,
		ConfigPath: configPath,
	}
}

func initEnvironArgs() {
	environArgs = &EnvironArgs{
		ConfigPath:       "",
		PYPath:           "",
		NetworkInterface: "",
	}
	pyPath, ok := os.LookupEnv("PY_LEXICAL_ANALYZER_PATH")
	if ok {
		environArgs.PYPath = pyPath
	}
	networkInterface, ok := os.LookupEnv("NETWORK_INTERFACE")
	if ok {
		environArgs.NetworkInterface = networkInterface
	}
}

func setConfByEnvirons() {
	// set env args, env variable has higher priority
	if environArgs.PYPath != "" {
		Conf.PythonLexicalAnalyzerPath = environArgs.PYPath
	}
	if environArgs.NetworkInterface != "" {
		Conf.NetworkInterface = environArgs.NetworkInterface
	}
}
