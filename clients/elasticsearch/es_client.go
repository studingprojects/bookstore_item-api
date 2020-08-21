package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/studingprojects/bookstore_utils-go/logger"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	log := logger.GetLogger()
	// Create a client
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200/"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetSniff(false),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)
	if err != nil {
		// Handle error
		panic(err)
	}
	// defer client.Stop()
	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		Type(docType).
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Get(index string, docType string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().
		Index(index).
		Type(docType).
		Id(id).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get id %s", id), err)
		return nil, err
	}
	return result, nil
}
