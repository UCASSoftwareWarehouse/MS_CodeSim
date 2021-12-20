package es

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestNewCodePlainText(t *testing.T) {
	c := NewCodePlainText("123", &ProjectFileIdentifier{
		CodeUniquePath: "asd",
		Tag:            "123",
		ID:             "idid",
		ProjectName: "some_project_name",
	})
	s, _ := json.Marshal(c)
	t.Log(string(s))

	f := "x.xx.go"
	i := strings.LastIndex(f, ".")
	if i == -1 {
		t.Error("i==-1")
	}
	_ = f[i+1:]
}
