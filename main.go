package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("helloworld"))
	})
	addr := "0.0.0.0:12138"
	log.Println("listen: ", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println(err)
	}
}
