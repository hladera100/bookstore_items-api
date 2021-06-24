package services

import (
	"net/http"

	"github.com/hladera100/bookstore_items-api/domain/items"
	"github.com/hladera100/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(item items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(items.Item) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me", http.StatusNotImplemented, "not_implemented", nil)
}

func (s *itemsService) Get(str string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me", http.StatusNotImplemented, "not_implemented", nil)
}
