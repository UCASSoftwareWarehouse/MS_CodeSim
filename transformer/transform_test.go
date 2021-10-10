package transformer

import (
	"encoding/json"
	"testing"
)

func TestTransformPythonCode(t *testing.T) {
	transformed, err := TransformPythonCode("/Users/purchaser/go/src/code_sim/transformer/python-lexical-analyzer/example.py")
	res, err := json.Marshal(transformed)
	t.Logf("transformed: [%s], err: [%v]", res, err)
}
