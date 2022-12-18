package main

import (
	"net/http"
)

//func main() {
//	http.HandleFunc("/hello", HelloHandler)
//	_ = http.ListenAndServe("localhost:8080", nil)
//}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello Go!"))
}
