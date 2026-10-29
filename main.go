package main

import (
 "encoding/json"
 "log"
 "net/http"
 "os"
 "strings"
)

func main() {
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Path
  if path == "/" { http.ServeFile(w, r, "index.html"); return }
  if strings.HasSuffix(path, ".html") || strings.HasSuffix(path, ".png") || strings.HasSuffix(path, ".pdf") || strings.HasSuffix(path, ".md") || strings.HasSuffix(path, ".txt") || strings.HasSuffix(path, ".json") {
   http.ServeFile(w, r, "."+path); return
  }
  http.ServeFile(w, r, "index.html")
 })
 http.HandleFunc("/api/balance", func(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type","application/json"); w.Header().Set("Access-Control-Allow-Origin","*")
  json.NewEncoder(w).Encode(map[string]interface{}{
   "address":"EKI_3774968e0ae21a32","balance":"21.000.000 EKI","validated":true,"hash":"70cd416462f3aef6","price":"R$0,01","p2p":":6001",
  })
 })
 http.HandleFunc("/api/buy", func(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin","*")
  json.NewEncoder(w).Encode(map[string]string{"status":"PIX Gerado","pix":"EKI_3774968e0ae21a32","whatsapp":"+55 17 9XXXX-XXXX"})
 })
 port:=os.Getenv("PORT"); if port==""{port="8080"}
 log.Fatal(http.ListenAndServe(":"+port, nil))
}
