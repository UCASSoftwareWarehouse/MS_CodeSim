package main

import (
	"code_sim/config"
	"code_sim/es"
	"code_sim/server"
)

func main() {
	// config.InitConfig()
	config.InitConfigDefault()
	es.InitEsCli()
	server.StartServe()
}

func matchTest() {
	//indices := []Document{
	//	NewCodePlainTextIndex("def some_func:\n\t\tpass", "/yzchnb/some_project/some_func.py", "v1.0"),
	//	NewCodePlainTextIndex("def some_other_func:\n\t\tpass", "/yzchnb/some_project/some_other_func.py", "v1.0"),
	//	NewCodePlainTextIndex("def another_func1:\n\t\tpass", "/yzchnb/some_project/another_func1.py", "v1.0"),
	//	NewCodePlainTextIndex("def another_func2:\n\t\tpass", "/yzchnb/some_project/another_func2.py", "v1.0"),
	//	NewCodePlainTextIndex("def another_func3:\n\t\tpass", "/yzchnb/some_project/another_func3.py", "v1.0"),
	//}
	// var indexName = "code-plain-text-index"
	// BulkIndex(indexName, indices)
	//ids := []string{
	//	indices[0].getID(), indices[1].getID(),
	//}
	// BulkDelete(indexName, ids)
	// matchPlainTest()
	//ID2Score := es.MatchPlain("def")
	//log.Printf("MatchPlain found ID2Score = %+v", ID2Score)
}
