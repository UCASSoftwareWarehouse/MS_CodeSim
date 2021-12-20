package es

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"log"
	"strings"
	"sync/atomic"
	"time"

	"github.com/dustin/go-humanize"
)

func BulkIndex(indexName IndexName, datas []Document) error {
	log.SetFlags(0)
	numWorkers := 5
	flushBytes := 5e+6
	var countSuccessful uint64

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         string(indexName), // The default index name
		Client:        ES,                // The Elasticsearch client
		NumWorkers:    numWorkers,        // The number of worker goroutines
		FlushBytes:    int(flushBytes),   // The flush threshold in bytes
		FlushInterval: 30 * time.Second,  // The periodic flush interval
	})
	if err != nil {
		log.Printf("Error creating the indexer: %s", err)
		return err
	}
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
	start := time.Now().UTC()

	// Loop over the collection
	//
	for _, a := range datas {
		// Prepare the data payload: encode article to JSON
		//
		data, err := json.Marshal(a)
		if err != nil {
			log.Printf("Cannot encode code-plain-text-index ID = [%s]: %s", a.getID(), err)
			return err
		}

		// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
		//
		// Add an item to the BulkIndexer
		//
		err = bi.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				// Action field configures the operation to perform (index, create, delete, update)
				Action: "index",

				// DocumentID is the (optional) document ID
				DocumentID: a.getID(),

				// Body is an `io.Reader` with the payload
				Body: bytes.NewReader(data),

				// OnSuccess is called for each successful operation
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					atomic.AddUint64(&countSuccessful, 1)
				},

				// OnFailure is called for each failed operation
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						log.Printf("ERROR: %s", err)
					} else {
						log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
		if err != nil {
			log.Printf("Unexpected error: %s", err)
			return err
		}
		// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
	}

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// Close the indexer
	//
	if err := bi.Close(context.Background()); err != nil {
		log.Printf("Unexpected error: %s", err)
		return err
	}
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

	biStats := bi.Stats()

	// Report the results: number of indexed docs, number of errors, duration, indexing rate
	//
	log.Println(strings.Repeat("▔", 65))

	dur := time.Since(start)

	if biStats.NumFailed > 0 {
		log.Printf(
			"Indexed [%s] documents with [%s] errors in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			humanize.Comma(int64(biStats.NumFailed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
		return fmt.Errorf("index has some failed attempts, numFailed=[%d]", biStats.NumFailed)
	} else {
		log.Printf(
			"Sucessfuly indexed [%s] documents in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
		return nil
	}
}

func BulkDeleteByTerms(indexName IndexName, terms map[string]string) error {
	// "query": {
	//    "bool": {
	//      "must": [
	//        {
	//          "term": { "project_name": "yzc_proj" }
	//        },
	//        {
	//          "term": { "tag": "1.0" }
	//        }
	//      ]
	//    }
	//  }
	must := make([]interface{}, 0)
	for termName, termValue := range terms {
		must = append(must, map[string]interface{}{
			"term": map[string]interface{}{
				termName: termValue,
			},
		})
	}
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": must,
			},
		},
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("BulkDeleteByTerms Fatal json marshal failed, err=[%s]", err)
		return err
	}
	res, err := ES.DeleteByQuery([]string{string(indexName)}, &buf)
	if err != nil {
		log.Printf("BulkDeleteByTerms DeleteByQuery failed, err=[%s]", err)
		return err
	}
	s, _ := json.Marshal(res)
	log.Printf("BulkDeleteByTerms DeleteByQuery response=[%s]", s)
	return nil
}

func BulkDelete(indexName string, ids []string) {
	log.SetFlags(0)
	numWorkers := 5
	flushBytes := 5e+6
	var countSuccessful uint64

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         indexName,        // The default index name
		Client:        ES,               // The Elasticsearch client
		NumWorkers:    numWorkers,       // The number of worker goroutines
		FlushBytes:    int(flushBytes),  // The flush threshold in bytes
		FlushInterval: 30 * time.Second, // The periodic flush interval
	})
	if err != nil {
		log.Fatalf("Error creating the indexer: %s", err)
	}
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
	start := time.Now().UTC()

	// Loop over the collection
	//
	for _, ID := range ids {
		// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
		//
		// Add an item to the BulkIndexer
		//
		err = bi.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				// Action field configures the operation to perform (index, create, delete, update)
				Action: "delete",

				// DocumentID is the (optional) document ID
				DocumentID: ID,

				// OnSuccess is called for each successful operation
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					atomic.AddUint64(&countSuccessful, 1)
				},

				// OnFailure is called for each failed operation
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						log.Printf("ERROR: %s", err)
					} else {
						log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
		if err != nil {
			log.Fatalf("Unexpected error: %s", err)
		}
		// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
	}

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// Close the indexer
	//
	if err := bi.Close(context.Background()); err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

	biStats := bi.Stats()

	// Report the results: number of indexed docs, number of errors, duration, indexing rate
	//
	log.Println(strings.Repeat("▔", 65))

	dur := time.Since(start)

	if biStats.NumFailed > 0 {
		log.Fatalf(
			"Delete [%s] documents with [%s] errors in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			humanize.Comma(int64(biStats.NumFailed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
	} else {
		log.Printf(
			"Sucessfuly delete [%s] documents in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
	}
}
