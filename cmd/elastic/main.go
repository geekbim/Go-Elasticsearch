package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/geekbim/Go-Elasticsearch/internal/pkg/storage/elasticsearch"
	"github.com/geekbim/Go-Elasticsearch/internal/post"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Bootstrap elasticsearch
	elastic, err := elasticsearch.New([]string{"http://localhost:9200"})
	if err != nil {
		log.Fatalln(err)
	}
	if err := elastic.CreateIndex("post"); err != nil {
		fmt.Println("ok")
		log.Fatalln(err)
	}

	// Bootstrap storage
	storage, err := elasticsearch.NewPostStorage(*elastic)
	if err != nil {
		log.Fatalln(err)
	}

	// Bootstrap API
	postAPI := post.New(storage)

	// Bootstrap HTTP router
	router := httprouter.New()
	router.HandlerFunc("POST", "/api/v1/posts", postAPI.Create)
	router.HandlerFunc("PATCH", "/api/v1/posts/:id", postAPI.Update)
	router.HandlerFunc("DELETE", "/api/v1/posts/:id", postAPI.Delete)
	router.HandlerFunc("GET", "/api/v1/posts/:id", postAPI.Find)

	// Start HTTP server
	log.Fatalln(http.ListenAndServe(":3000", router))
}
