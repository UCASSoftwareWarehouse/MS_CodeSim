package es

import (
	"code_sim/config"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

var ES *elasticsearch.Client

func InitEsCli() {
	cfg := elasticsearch.Config{
		Addresses: []string{config.Conf.ESAddr},
	}
	esCli, err := elasticsearch.NewClient(cfg)
	// esCli, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	ES = esCli
}
