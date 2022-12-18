package main

import (
	"net/http"
	"sync"
)

func StartGinApi() {
	r := setUpRouter()
	_ = r.Run("localhost:8080")
}

func StartHttpApi() {
	http.HandleFunc("/hello", HelloHandler)
	_ = http.ListenAndServe("localhost:8000", nil)
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		StartHttpApi()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		StartGinApi()
		wg.Done()
	}()

	wg.Wait()
}
