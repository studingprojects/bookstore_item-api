package items

import (
	"errors"

	"github.com/studingprojects/bookstore_item-api/clients/elasticsearch"
	"github.com/studingprojects/bookstore_utils-go/rest_errors"
)

var (
	itemIndex = "items"
)

func (i *Item) Save() *rest_errors.RestErr {
	_, err := elasticsearch.Client.Index(itemIndex, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	return nil
}
