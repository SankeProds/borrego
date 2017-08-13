package main

import (
	"fmt"
	"net/http"
	"github.com/SankeProds/borrego/handlers"
)

func main() {
	fmt.Printf("BorreGO!\n")
	http.Handle("/test", http.HandlerFunc(handlers.Test))
	http.ListenAndServe(":8080", nil)
}
