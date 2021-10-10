package update

import (
	"code_sim/config"
	"code_sim/es"
	"code_sim/internal/converter"
	"code_sim/pb_gen"
	"path/filepath"
	"testing"
)

func TestUploader_DoUpload(t *testing.T) {
	res, err := filepath.Rel("/abc", "abc")
	t.Logf("res=[%v], err=[%v]", res, err)
}

func TestDocGenerator(t *testing.T) {
	config.InitConfigWithFile("/Users/purchaser/go/src/code_sim/config.yml", config.DevEnv)
	gen := docGenerators[es.CodeTransformedTextIndex]
	relP := "some_path/some_file.py"
	doc := gen(converter.ConvertToES(&pb_gen.CodeSimProjectFile{
		ProjectInfo: &pb_gen.CodeSimProject{
			ProjectName: "some_proj",
			Tag:         "v1.0",
		},
		RelativePath: relP,
	}), "iavgc= asr431 z<p ===a1 ;;a' \"main\":", relP)
	t.Logf("doc=[%+v]", doc)
}
