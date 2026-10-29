package main

import (
 "crypto/rand"
 "encoding/hex"
 "encoding/json"
 "log"
 "net/http"
 "os"
)
func randomHex(n int) string {
 b := make([]byte, n)
 rand.Read(b)
 return hex.EncodeToString(b)
}
func main(){
 port := os.Getenv("PORT")
 if port==""{ port="10000" }
 http.HandleFunc("/wallet/create", func(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type","application/json")
  w.Header().Set("Access-Control-Allow-Origin","*")
  priv := randomHex(32)
  json.NewEncoder(w).Encode(map[string]string{
   "address": "eki1"+priv[:38],
   "mnemonic": "salve esta chave "+randomHex(8),
   "privateKey": priv,
  })
 })
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
  if r.URL.Path=="/" { http.ServeFile(w, r, "index.html"); return }
  if r.URL.Path=="/logo.jpg" { http.ServeFile(w, r, "logo.jpg"); return }
  http.NotFound(w,r)
 })
 log.Println("EKI Mainnet OK on :"+port)
 log.Fatal(http.ListenAndServe(":"+port, nil))
}
