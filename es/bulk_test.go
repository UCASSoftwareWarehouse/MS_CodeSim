package es

import (
	"math/rand"
	"strings"
)

//
//func TestBulkIndex(t *testing.T) {
//	config.InitConfig()
//	InitEsCli()
//	docs := make([]Document, 0)
//	for i := 0; i < 10; i++ {
//		p := &pb_gen.CodeSimProject{
//			ProjectName: "yzchnb/SomeProject",
//			Tag:         "v1.0",
//		}
//		relP := fmt.Sprintf("some_dir/%d", i)
//		pt := fmt.Sprintf("plainText:%s:::::%d", genLongStr(), i)
//		esInfo :=
//		docs = append(docs, NewCodePlainText(pt, esInfo))
//	}
//	err := BulkIndex(CodePlainTextIndex, docs)
//	t.Logf("err=[%v]", err)
//}

func genLongStr() string {
	sb := &strings.Builder{}
	sb.Grow(1000)
	for i := 0; i < 10000; i++ {
		sb.WriteByte(byte(rand.Intn(128)))
	}
	return sb.String()
}
