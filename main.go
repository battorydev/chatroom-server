package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World from Go!")
}