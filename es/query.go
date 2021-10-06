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

type Search struct {
	Query *Match    `json:"query"`
	WithSource bool `json:"_source"`
	Sort []SortItem `json:"sort"`
	From int `json:"from"`
	Size int `json:"size"`
}

type Match struct {
	M map[string]*MatchOpt `json:"match"` // fieldName to MatchOpt
}

type MatchOpt struct {
	QueryText string `json:"query"`
	Analyzer string `json:"analyzer"`
}

type SortItem map[string]string

func buildSortItem(field string, desc bool) SortItem {
	direction := "asc"
	if desc {
		direction = "desc"
	}
	return SortItem{field: direction}
}

var indexName2CodeFieldName = map[IndexName]string{
	CodePlainTextIndex: "code-plain-text",
	CodeTransformedTextIndex: "code-transformed-text",
}

func MatchCode(targetCode string, targetIndexName IndexName, from, size int) (ID2Score map[string]float64) {
	indexName := targetIndexName
	if _, ok := indexName2CodeFieldName[indexName]; !ok {
		log.Printf("Unsupportted indexName = [%s]", indexName)
		return nil
	}
	var buf bytes.Buffer
	search := &Search{
		Query:&Match{
			map[string]*MatchOpt{
				indexName2CodeFieldName[targetIndexName]: {
					QueryText: targetCode,
					Analyzer:  "plain_text_ngram_analyzer",
				},
			},
		},
		WithSource: false,
		Sort: []SortItem{buildSortItem("_score", true)},
		From: from,
		Size: size,
	}
	if err := json.NewEncoder(&buf).Encode(search); err != nil {
		log.Printf("Error encoding query: %s", err)
		return nil
	}
	log.Printf("search query payload: %s", buf.String())

	res, err := ES.Search(
		ES.Search.WithContext(context.Background()),
		ES.Search.WithIndex(string(indexName)),
		ES.Search.WithBody(&buf),
		ES.Search.WithTrackTotalHits(true),
		ES.Search.WithPretty(),
	)

	if err != nil {
		log.Printf("Error getting response: %s", err)
		return nil
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
		return nil
	}

	log.Printf("res to String = [%s]", res.String())

	var r map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		return nil
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
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
