package main

import (
    "fmt"
    "net/http"
    "log"
)

type Schema struct {
    author string
}

func IndexHandler (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Helo World!")
}

func main() {
    http.HandleFunc("/", IndexHandler)
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}