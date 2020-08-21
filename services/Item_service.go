package services

import (
	"github.com/studingprojects/bookstore_item-api/domain/items"
	"github.com/studingprojects/bookstore_utils-go/rest_errors"
)

var (
	ItemService itemServiceIterface = &itemService{}
)

type itemServiceIterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemService struct{}

func (s *itemService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Get(id string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{Id: id}
	if err := item.Get(); err != nil {
		return nil, err
	}

	return &item, nil
}
