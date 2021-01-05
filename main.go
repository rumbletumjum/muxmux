package main

import (
	"fmt"
	"log"
	"net/http"

	mux "muxmux.rkbkr.io/router"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"chillin"}`))
}

func postRoot(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["msg"]
	if !ok || len(keys[0]) < 1 {
		log.Println("URL Param 'msg' is missing")
		return
	}

	log.Printf("Got params: %v", keys)
	msg := keys[0]
	res := fmt.Sprintf(`{"your_message":"%s"}`, msg)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func main() {
	router := mux.NewRouter()

	router.GET("/", getRoot)
	router.POST("/", postRoot)

	log.Fatal(http.ListenAndServe(":8888", router))
}
