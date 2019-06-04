package main

import (
  "fmt"
  "net/http"
  "time"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello Amazing Go!")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Awesome About")
}

func main() {
  fmt.Println("Starting Web Application", time.Now())
  http.HandleFunc("/", index_handler)
  http.HandleFunc("/about", about_handler)
  fmt.Println("Web Application started", time.Now())

  fmt.Println("Status:", "listen and serve with port 8000", time.Now())
  http.ListenAndServe(":8000", nil)
}
