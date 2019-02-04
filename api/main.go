package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dontomato/go_web_api/mypkg"
)

func main() {
	router := NewRouter()

	var e = mypkg.TestModel{ID: 0, Name: "Jopa"}
	var r = mypkg.Te()
	fmt.Println(e, r)

	log.Fatal(http.ListenAndServe(":8283", router))
}
