package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hello go")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":9000", nil)
}