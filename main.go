package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("GET /", rootHandler)
	r.HandleFunc("GET /list", listHandler)

	log.Fatal(http.ListenAndServe(":9190", r))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Random number is : %v", GetRandomNumber())
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "사실 아무것도 없지롱")
}

func GetRandomNumber() int {
	return rand.Intn(1000)
}
