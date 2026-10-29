package main

import (
 "encoding/json"
 "fmt"
 "log"
 "net/http"
 "os"
)

func main() {
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path == "/" {
   http.ServeFile(w, r, "index.html")
   return
  }
  http.ServeFile(w, r, "."+r.URL.Path)
 })
 http.HandleFunc("/api/balance", func(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  json.NewEncoder(w).Encode(map[string]interface{}{
   "address": "EKI_3774968e0ae21a32",
   "balance": "21.000.000 EKI",
   "validated": true,
   "hash": "70cd416462f3aef6acfd9e47d9317cc6",
   "R": "f22c7960858b0d81704c",
   "S": "d80b3c38ca595a396712",
   "price": "R$0,01",
   "marketcap": "R$210.000,00",
   "p2p": ":6001",
   "api": ":8080",
  })
 })
 http.HandleFunc("/api/buy", func(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  json.NewEncoder(w).Encode(map[string]string{"status": "PIX QR gerado", "address": "EKI_3774968e0ae21a32"})
 })

 port := os.Getenv("PORT")
 if port == "" {
  port = "8080"
 }
 fmt.Printf("🚀 EKI Chain rodando! API :%s | P2P :6001\n", port)
 fmt.Printf("💰 TOKEN EKI NATIVO - BASE EKI | 1 EKI = R$0,01\n")
 log.Fatal(http.ListenAndServe(":"+port, nil))
}
