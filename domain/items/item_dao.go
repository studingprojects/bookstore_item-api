package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/studingprojects/bookstore_item-api/clients/elasticsearch"
	"github.com/studingprojects/bookstore_utils-go/rest_errors"
)

var (
	indexItem = "items"
	typeItem  = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItem, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItem, typeItem, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFounfError(fmt.Sprintf("item %s not found", i.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when to try get id %s", i.Id), errors.New("database error"))
	}
	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("error when to try parse result from database", err)
	}
	if parseErr := json.Unmarshal(bytes, i); parseErr != nil {
		return rest_errors.NewInternalServerError("error when to try parse result from database", parseErr)
	}
	i.Id = itemId
	return nil
}
