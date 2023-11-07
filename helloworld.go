package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", MyHandler1)
	http.HandleFunc("/john", MyHandler2)
	http.ListenAndServe(":8081", nil)
}
func MyHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}
func MyHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello John\n")
}
