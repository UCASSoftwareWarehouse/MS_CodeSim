package es

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strings"
)

type Query struct{}

func GetQuery() Query {
	return Query{}
}

type MatchCodeParams struct {
	TargetCode   string
	TargetFields []CodeIndexField
	From         int
	Size         int
	WithSource   bool
}

type MatchCodeEachRes struct {
	ID        string
	Score     float64
	PlainText string
}

// MatchCode
// GET test-index/_search
//{
//  "query": {
//    "multi_match": {
//      "query": "import",
//      "fields": ["code-plain-text.golang"],
//      "type": "most_fields"
//    }
//  },
//  "from": 0,
//  "size": 10
//}
func (q Query) MatchCode(params *MatchCodeParams) []*MatchCodeEachRes {
	type MatchOpt struct {
		QueryText string   `json:"query"`
		Fields    []string `json:"fields"`
		Type      string   `json:"type"`
	}
	type MultiMatch struct {
		M *MatchOpt `json:"multi_match"`
	}
	type SortItem map[string]string
	type Search struct {
		Query      *MultiMatch `json:"query"`
		WithSource bool        `json:"_source"`
		Sort       []SortItem  `json:"sort"`
		From       int         `json:"from"`
		Size       int         `json:"size"`
	}
	buildSortItem := func(field string, desc bool) SortItem {
		direction := "asc"
		if desc {
			direction = "desc"
		}
		return SortItem{field: direction}
	}
	var buf bytes.Buffer
	getFields := func() []string {
		res := make([]string, 0, len(params.TargetFields))
		for _, f := range params.TargetFields {
			res = append(res, string(f))
		}
		return res
	}
	search := &Search{
		Query: &MultiMatch{
			M: &MatchOpt{
				QueryText: params.TargetCode,
				Fields:    getFields(),
				Type:      "most_fields",
			},
		},
		WithSource: params.WithSource,
		Sort:       []SortItem{buildSortItem("_score", true)},
		From:       params.From,
		Size:       params.Size,
	}
	if err := json.NewEncoder(&buf).Encode(search); err != nil {
		log.Printf("Error encoding query: %s", err)
		return nil
	}
	log.Printf("search query payload: %s", buf.String())

	r, err := q.doQuery(CodeIndex, buf)
	if err != nil {
		return nil
	}

	// Print the ID and document source for each hit.
	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})
	res := make([]*MatchCodeEachRes, 0, len(hits))
	for _, hit := range hits {
		ID := hit.(map[string]interface{})["_id"].(string)
		log.Printf(" * ID=%s, %s", ID, hit.(map[string]interface{})["_source"])
		score := hit.(map[string]interface{})["_score"].(float64)
		var plainText string
		if params.WithSource {
			plainText = hit.(map[string]interface{})["_source"].(map[string]interface{})["code-plain-text"].(string)
		}
		res = append(res, &MatchCodeEachRes{
			ID:        ID,
			Score:     score,
			PlainText: plainText,
		})
	}
	log.Println(strings.Repeat("=", 37))
	return res
}

func (q Query) GetCodeByIDs(targetIDs []string) (map[string][]byte, error) {
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
	r, err := q.doQuery(CodeIndex, buf)
	if err != nil {
		return nil, err
	}
	// Print the ID and document source for each hit.
	hitID2Code := make(map[string][]byte)
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		ID := hit.(map[string]interface{})["_id"].(string)
		code := hit.(map[string]interface{})["_source"].(map[string]interface{})["code-plain-text"].(string)
		hitID2Code[ID] = []byte(code)
		log.Printf(" * ID=%s, code=%s", ID, code)
	}
	return hitID2Code, nil
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
