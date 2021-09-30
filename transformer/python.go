package transformer

import (
	"code_sim/config"
	"encoding/json"
	"log"
	"os/exec"
	"strings"
)

func TransformPythonCode(codeFilepath string) (string, error) {
	cmd := exec.Command("python", config.Conf.PythonLexicalAnalyzerPath, codeFilepath)
	buf, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("python lexicial analyzer execution was failed, err = [%s]", err)
		return "", err
	}
	s := string(buf)
	lines := strings.Split(s, "\n")
	transformed := strings.Join(lines, config.Conf.TransformCodeSplitter)
	marshaled, err := json.Marshal(transformed)
	if err != nil {
		return "", err
	}
	return string(marshaled), nil
}