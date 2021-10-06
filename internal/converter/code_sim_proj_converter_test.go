package converter

import (
	"testing"
)

func TestExtractProjectFromESID(t *testing.T) {
	info, err := ExtractProjectFileFromESID("yzchnb/SomeProj@aaa/bbb/ccc:v1.0")
	t.Logf("info=[%+v], err=[%+v]", info, err)
}
