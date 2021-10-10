package query

import "code_sim/es"

func GetCodeByIDs(esIDs []string) (map[string][]byte, error) {
	return es.GetQuery().FindCodeByIDs(es.CodePlainTextIndex, esIDs)
}