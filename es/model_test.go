package es

import (
	"code_sim/transformer"
	"encoding/json"
	"strings"
	"testing"
)

func TestNewCodePlainText(t *testing.T) {
	c := NewCodePlainText("123", &ProjectFileIdentifier{
		CodeUniquePath: "asd",
		Tag:            "123",
		ID:             "idid",
	})
	s, _ := json.Marshal(c)
	t.Log(string(s))

	f := "x.xx.go"
	i := strings.LastIndex(f, ".")
	if i == -1 {
		t.Error("i==-1")
	}
	suffix := f[i+1:]
	codeType, err := transformer.GetSupportedCodeType(suffix)
	t.Logf("codeType=[%v], err=[%v]", codeType, err)
}
