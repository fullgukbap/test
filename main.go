package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", RootHandler)

	log.Fatal(http.ListenAndServe(":9000", r))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Random number is : %v", GetRandomNumber())
}

func GetRandomNumber() int {
	return rand.Intn(1000)
}
