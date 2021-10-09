package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ConfigurationEnv string

const (
	DevEnv ConfigurationEnv = "dev"
	PrdEnv ConfigurationEnv = "prd"
)

type Configuration map[ConfigurationEnv]*EachConfig

type EachConfig struct {
	AppName                   string `yaml:"app_name"`
	Host                      string `yaml:"host"`
	Port                      int    `yaml:"port"`
	ESAddr                    string `yaml:"es_addr"`
	PythonLexicalAnalyzerPath string `yaml:"python_lexical_analyzer_path"`
	TransformCodeSplitter     string `yaml:"transform_code_splitter"`
	ConsulAddr                string `yaml:"consul_addr"`
	NetworkInterface          string `yaml:"network_interface"`

	Env ConfigurationEnv
}

type CmdArgs struct {
	Env        ConfigurationEnv
	Port       int
	ConfigPath string
}

type EnvironArgs struct {
	ConfigPath       string
	PYPath           string
	NetworkInterface string
}

func (c *EachConfig) GetEnv() ConfigurationEnv {
	return c.Env
}

func parse() Configuration {
	configPath := cmdArgs.ConfigPath
	if cmdArgs.ConfigPath == "" {
		if environArgs.ConfigPath == "" {
			panic("CONFIG_PATH not set, plz check your environs")
		}
		configPath = environArgs.ConfigPath
	}
	bs, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("ConfigForEnv parse failed, read file failed, err=[%v]", err)
	}
	conf := make(Configuration)
	err = yaml.Unmarshal(bs, &conf)
	if err != nil {
		log.Fatalf("ConfigForEnv parse failed, unmarshal config failed, err=[%v]", err)
	}
	return conf
}
