package es

import (
	"code_sim/config"
	"testing"
)

func TestMatchCode(t *testing.T) {
	config.InitConfig()
	InitEsCli()
	res := MatchCode("c", CodePlainTextIndex, 0, 100)
	t.Logf("res=[%v]", res)
}
