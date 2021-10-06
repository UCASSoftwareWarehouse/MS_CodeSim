package data

import (
	"path/filepath"
	"testing"
)

func TestUploader_DoUpload(t *testing.T) {
	res, err := filepath.Rel("/abc", "abc")
	t.Logf("res=[%v], err=[%v]", res, err)
}
