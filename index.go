package main

import (
        "net/http"
)

func main() {
        http.Handle("/", http.FileServer(http.Dir("./index.html")))
        http.ListenAndServe(":80", nil)
}