//hello (web) server
package main

import (
  "fmt"
  "net/http"
  "html/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("index.html")
  t.Execute(w, r.URL)
}

func main() {
  fmt.Println("Starting server...")

  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
