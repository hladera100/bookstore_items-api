package controllers

import (
	"fmt"
	"net/http"

	"github.com/federicoleon/bookstore_oauth-go/oauth"
	"github.com/hladera100/bookstore_items-api/domain/items"
	"github.com/hladera100/bookstore_items-api/services"
)

var (
	ItemsControllers itemsControllersInterface = &itemsController{}
)

type itemsControllersInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//todo return error json to the caller
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		//todo return error json to the caller
	}
	fmt.Println(result)
	//TODO: Return created  item with HTTP status 201 -Created
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
