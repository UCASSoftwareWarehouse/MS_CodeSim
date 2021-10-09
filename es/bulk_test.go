package es

import (
	"code_sim/config"
	"code_sim/internal/converter"
	"code_sim/pb_gen"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func TestBulkIndex(t *testing.T) {
	config.InitConfigDefault()
	InitEsCli()
	docs := make([]Document, 0)
	for i := 0; i < 10; i++ {
		p := &pb_gen.CodeSimProject{
			ProjectName: "yzchnb/SomeProject",
			Tag:         "v1.0",
		}
		relP := fmt.Sprintf("some_dir/%d", i)
		pt := fmt.Sprintf("plainText:%s:::::%d", genLongStr(), i)
		cup, tag, ID := converter.GenEsInfo(&pb_gen.CodeSimProjectFile{
			ProjectInfo:  p,
			RelativePath: relP,
		})
		docs = append(docs, NewCodePlainText(pt, cup, tag, ID))
	}
	err := BulkIndex(CodePlainTextIndex, docs)
	t.Logf("err=[%v]", err)
}

func genLongStr() string {
	sb := &strings.Builder{}
	sb.Grow(1000)
	for i := 0; i < 10000; i++ {
		sb.WriteByte(byte(rand.Intn(128)))
	}
	return sb.String()
}
