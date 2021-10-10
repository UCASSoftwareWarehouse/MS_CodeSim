package es

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strings"
)

type IndexName string

const (
	CodePlainTextIndex       IndexName = "code-plain-text-index"
	CodeTransformedTextIndex IndexName = "code-transformed-text-index"
)

var indexName2CodeFieldName = map[IndexName]string{
	CodePlainTextIndex:       "code-plain-text",
	CodeTransformedTextIndex: "code-transformed-text",
}

var indexName2CodeFileAnalyzer = map[IndexName]string{
	CodePlainTextIndex:       "plain_text_ngram_analyzer",
	CodeTransformedTextIndex: "transformed_text_analyzer",
}

type Query struct{}

func GetQuery() Query{
	return Query{}
}

func (q Query) MatchCodeIDs(targetCode string, targetIndexName IndexName, from, size int) (ID2Score map[string]float64) {
	type MatchOpt struct {
		QueryText string `json:"query"`
		Analyzer  string `json:"analyzer"`
	}
	type Match struct {
		M map[string]*MatchOpt `json:"match"` // fieldName to MatchOpt
	}
	type SortItem map[string]string
	type Search struct {
		Query      *Match     `json:"query"`
		WithSource bool       `json:"_source"`
		Sort       []SortItem `json:"sort"`
		From       int        `json:"from"`
		Size       int        `json:"size"`
	}
	buildSortItem := func(field string, desc bool) SortItem {
		direction := "asc"
		if desc {
			direction = "desc"
		}
		return SortItem{field: direction}
	}
	var buf bytes.Buffer
	search := &Search{
		Query: &Match{
			map[string]*MatchOpt{
				indexName2CodeFieldName[targetIndexName]: {
					QueryText: targetCode,
					Analyzer:  indexName2CodeFileAnalyzer[targetIndexName],
				},
			},
		},
		WithSource: false,
		Sort:       []SortItem{buildSortItem("_score", true)},
		From:       from,
		Size:       size,
	}
	if err := json.NewEncoder(&buf).Encode(search); err != nil {
		log.Printf("Error encoding query: %s", err)
		return nil
	}
	log.Printf("search query payload: %s", buf.String())

	r, err := q.doQuery(targetIndexName, buf)
	if err != nil {
		return nil
	}

	// Print the ID and document source for each hit.
	hitID2Scores := make(map[string]float64)
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		ID := hit.(map[string]interface{})["_id"].(string)
		hitID2Scores[ID] = hit.(map[string]interface{})["_score"].(float64)
		log.Printf(" * ID=%s, %s", ID, hit.(map[string]interface{})["_source"])
	}
	log.Println(strings.Repeat("=", 37))
	return hitID2Scores
}

func (q Query) doQuery(indexName IndexName, buf bytes.Buffer) (map[string]interface{}, error) {
	res, err := ES.Search(
		ES.Search.WithContext(context.Background()),
		ES.Search.WithIndex(string(indexName)),
		ES.Search.WithBody(&buf),
		ES.Search.WithTrackTotalHits(true),
		ES.Search.WithPretty(),
	)

	if err != nil {
		log.Printf("Error getting response: %s", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
		return nil, err
	}
	log.Printf("res to String = [%s]", res.String())
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		return nil, err
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	return r, nil
}

func (q Query) FindCodeByIDs(targetIndexName IndexName, targetIDs []string) (map[string][]byte, error) {
	type IDs struct {
		Values []string `json:"values"`
	}
	type Match struct {
		IDs *IDs `json:"ids"`
	}
	type FindByIDs struct {
		Query      *Match `json:"query"`
		WithSource bool   `json:"_source"`
	}
	var buf bytes.Buffer
	findByIDs := &FindByIDs{
		Query: &Match{
			IDs: &IDs{
				Values: targetIDs,
			},
		},
		WithSource: true,
	}
	if err := json.NewEncoder(&buf).Encode(findByIDs); err != nil {
		log.Printf("Error encoding query: %s", err)
		return nil, err
	}
	log.Printf("find by ids query payload: %s", buf.String())
	r, err := q.doQuery(targetIndexName, buf)
	if err != nil {
		return nil, err
	}
	// Print the ID and document source for each hit.
	hitID2Code := make(map[string][]byte)
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		ID := hit.(map[string]interface{})["_id"].(string)
		code := hit.(map[string]interface{})["_source"].(map[string]interface{})[indexName2CodeFieldName[targetIndexName]].(string)
		hitID2Code[ID] = []byte(code)
		log.Printf(" * ID=%s, code=%s", ID, code)
	}
	return hitID2Code, nil
}

//
//func MatchPlain(targetPlainCode string) (ID2Score map[string]float64) {
//	indexName := CodePlainTextIndex
//	var buf bytes.Buffer
//	search := &Search{
//		Query:&Match{
//			map[string]*MatchOpt{
//				indexName2CodeFieldName[indexName]: {
//					QueryText: targetPlainCode,
//					Analyzer:  "plain_text_ngram_analyzer",
//				},
//			},
//		},
//		WithSource: false,
//		Sort: []SortItem{buildSortItem("_score", true)},
//	}
//	if err := json.NewEncoder(&buf).Encode(search); err != nil {
//		log.Fatalf("Error encoding query: %s", err)
//	}
//	log.Printf("search query payload: %s", buf.String())
//
//	res, err := ES.Search(
//		ES.Search.WithContext(context.Background()),
//		ES.Search.WithIndex(indexName),
//		ES.Search.WithBody(&buf),
//		ES.Search.WithTrackTotalHits(true),
//		ES.Search.WithPretty(),
//	)
//
//	if err != nil {
//		log.Fatalf("Error getting response: %s", err)
//	}
//	defer res.Body.Close()
//
//	if res.IsError() {
//		var e map[string]interface{}
//		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
//			log.Fatalf("Error parsing the response body: %s", err)
//		} else {
//			// Print the response status and error information.
//			log.Fatalf("[%s] %s: %s",
//				res.Status(),
//				e["error"].(map[string]interface{})["type"],
//				e["error"].(map[string]interface{})["reason"],
//			)
//		}
//	}
//
//	log.Printf("res to String = [%s]", res.String())
//
//	var r map[string]interface{}
//
//	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
//		log.Fatalf("Error parsing the response body: %s", err)
//	}
//	// Print the response status, number of results, and request duration.
//	log.Printf(
//		"[%s] %d hits; took: %dms",
//		res.Status(),
//		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
//		int(r["took"].(float64)),
//	)
//	// Print the ID and document source for each hit.
//	hitID2Scores := make(map[string]float64)
//	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
//		ID := hit.(map[string]interface{})["_id"].(string)
//		hitID2Scores[ID] = hit.(map[string]interface{})["_score"].(float64)
//		log.Printf(" * ID=%s, %s", ID, hit.(map[string]interface{})["_source"])
//	}
//	log.Println(strings.Repeat("=", 37))
//	return hitID2Scores
//}
