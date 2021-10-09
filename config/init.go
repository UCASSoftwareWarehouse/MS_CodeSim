package config

import "os"

var Conf *EachConfig

func InitConfig(configFilepath string, env ConfigurationEnv) {
	c := parse(configFilepath)
	Conf = c[env]
	setByEnvirons()
}

func setByEnvirons() {
	pyPath, ok := os.LookupEnv("PY_LEXICAL_ANALYZER_PATH")
	if ok {
		Conf.PythonLexicalAnalyzerPath = pyPath
	}
}
