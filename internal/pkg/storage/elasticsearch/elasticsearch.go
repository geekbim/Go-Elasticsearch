package elasticsearch

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
)

type Elasticsearch struct {
	client *elasticsearch.Client
	index  string
	alias  string
}

func New(addresses []string) (*Elasticsearch, error) {
	cfg := elasticsearch.Config{
		Addresses: addresses,
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &Elasticsearch{
		client: client,
	}, nil
}

func (e *Elasticsearch) CreateIndex(index string) error {
	e.index = index
	e.alias = index + "_alias"

	res, err := e.client.Indices.Exists([]string{e.index})
	if err != nil {
		return fmt.Errorf("cannot check index existence: %w", err)
	}
	if res.StatusCode == http.StatusOK {
		return nil
	}
	if res.StatusCode != http.StatusBadRequest {
		return fmt.Errorf("error in index existence response: %s", res.String())
	}

	res, err = e.client.Indices.Create(e.index)
	if err != nil {
		return fmt.Errorf("cannot create index: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("error in index creation response: %s", res.String())
	}

	res, err = e.client.Indices.PutAlias([]string{e.index}, e.alias)
	if err != nil {
		return fmt.Errorf("cannot create index alias: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("error in index alias creation response: %s", res.String())
	}

	return nil
}

// document represents a single document in Get API response body.
type document struct {
	Source interface{} `json:"_source"`
}
