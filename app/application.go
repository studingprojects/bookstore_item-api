package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/studingprojects/bookstore_item-api/clients/elasticsearch"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	mapURLs()
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8087",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

	fmt.Println("Start application ::8087")
}
