package main

import (
 "encoding/json"
 "fmt"
 "log"
 "net/http"
 "os"
)

func main() {
 port := os.Getenv("PORT")
 if port == "" {
  port = "10000"
 }
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Eki Chain Live - commit 8beae7e")
 })
 http.HandleFunc("/supply", func(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(map[string]interface{}{
   "supply": 1000000000,
   "symbol": "EKI",
  })
 })
 log.Printf("Starting on port %s", port)
 log.Fatal(http.ListenAndServe(":"+port, nil))
}
