package es6

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"
	"github.com/pkg/errors"
	"github.com/ugent-library/people/models"
)

type Config struct {
	ClientConfig   elasticsearch.Config
	Index          string
	Settings       string
	IndexRetention int // -1: keep all old indexes, >=0: keep x old indexes
}

type Client struct {
	Config
	es *elasticsearch.Client
}

type es6SearchReq struct {
	From  int        `json:"from"`
	Size  int        `json:"size"`
	Query models.M   `json:"query"`
	Sort  []models.M `json:"sort,omitempty"`
}
type es6SearchRes struct {
	Hits struct {
		Total int `json:"total"`
		Hits  []struct {
			Source json.RawMessage `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func NewClient(config Config) (*Client, error) {
	client, err := elasticsearch.NewClient(config.ClientConfig)
	if err != nil {
		return nil, err
	}
	return &Client{Config: config, es: client}, nil
}

func (c *Client) searchWithOpts(opts []func(*esapi.SearchRequest), responseBody any) error {

	res, err := c.es.Search(opts...)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		buf := &bytes.Buffer{}
		if _, err := io.Copy(buf, res.Body); err != nil {
			return err
		}
		return errors.New("Es6 error response: " + buf.String())
	}

	if err := json.NewDecoder(res.Body).Decode(responseBody); err != nil {
		return errors.Wrap(err, "Error parsing the response body")
	}

	return nil
}

func (c *Client) Search(req *es6SearchReq, responseBody any) error {

	/*
		IMPORTANT: do not use WithSort (which requires syntax like <field>:<direction>)
		Put "sort" in request body
	*/
	opts := []func(*esapi.SearchRequest){
		c.es.Search.WithContext(context.Background()),
		c.es.Search.WithIndex(c.Index),
		c.es.Search.WithTrackTotalHits(true),
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(req); err != nil {
		return err
	}
	opts = append(opts, c.es.Search.WithBody(&buf))

	return c.searchWithOpts(opts, responseBody)
}
