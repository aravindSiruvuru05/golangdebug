package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err == nil {
		log.Fatal("Server failed to start: ", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	sum := 5 + 6
	fmt.Fprintf(w, "sum of 5, 6 is %d", sum)
}
