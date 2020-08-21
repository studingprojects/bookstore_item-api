package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/studingprojects/bookstore_utils-go/rest_errors"

	"github.com/studingprojects/bookstore_item-api/utils/http_utils"

	"github.com/studingprojects/bookstore_item-api/domain/items"
	"github.com/studingprojects/bookstore_item-api/services"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: authentiate request HERE
	// if err := oauth.AuthenticateRequest(r); err != nil {
	// 	http_utils.ResponseJson(w, err.Status, err)
	// 	return
	// }

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.ResponseJson(w, restErr.Status(), restErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err = json.Unmarshal(requestBody, &itemRequest); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid body json")
		http_utils.ResponseJson(w, restErr.Status(), restErr)
		return
	}

	// TODO: assign seller id here
	// itemRequest.Seller = oauth.GetClientId(r)
	result, createErr := services.ItemService.Create(itemRequest)
	if createErr != nil {
		http_utils.ResponseJson(w, createErr.Status(), createErr)
		return
	}
	http_utils.ResponseJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	item, err := services.ItemService.Get(vars["id"])
	if err != nil {
		http_utils.ResponseJson(w, err.Status(), err)
		return
	}
	http_utils.ResponseJson(w, http.StatusOK, item)
}
