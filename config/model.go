package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ConfigurationEnv string

const (
	DevEnv ConfigurationEnv = "dev"
	PrdEnv ConfigurationEnv = "prd"
)

func convert2Env(env string) (ConfigurationEnv, error) {
	switch env {
	case "dev":
		return DevEnv, nil
	case "prd":
		return PrdEnv, nil
	default:
		return "", errors.New("illegal env, should be dev or prd")
	}
}

type Configuration map[ConfigurationEnv]*EachConfig

type EachConfig struct {
	AppName               string `yaml:"app_name"`
	Host                  string `yaml:"host"`
	Port                  int    `yaml:"port"`
	ESAddr                string `yaml:"es_addr"`
	TransformCodeSplitter string `yaml:"transform_code_splitter"`
	ConsulAddr            string `yaml:"consul_addr"`
	NetworkInterface      string `yaml:"network_interface"`
	MaxSearchDepth        int    `yaml:"max_search_depth"`

	Env ConfigurationEnv
}

type args struct {
	ConfigPath       string
	Env              ConfigurationEnv
	NetworkInterface string
	Port             int
}

func (c *EachConfig) GetEnv() ConfigurationEnv {
	return c.Env
}

func parse(args *args) *EachConfig {
	if args.Env == "" {
		panic("ENV not set, plz check your environs or cmd args")
	}
	if args.ConfigPath == "" {
		panic("CONFIG_PATH not set, plz check your environs or cmd args")
	}
	bs, err := ioutil.ReadFile(args.ConfigPath)
	if err != nil {
		log.Fatalf("ConfigForEnv parse failed, read file failed, err=[%v]", err)
	}
	configs := make(Configuration)
	err = yaml.Unmarshal(bs, &configs)
	if err != nil {
		log.Fatalf("ConfigForEnv parse failed, unmarshal config failed, err=[%v]", err)
	}

	conf := configs[args.Env]
	// reset settings by args
	conf.Env = args.Env
	if args.Port != 0 {
		conf.Port = args.Port
	}
	if args.NetworkInterface != "" {
		conf.NetworkInterface = args.NetworkInterface
	}

	return conf
}
