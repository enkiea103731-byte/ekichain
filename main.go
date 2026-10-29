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
  addr := "eki1"+priv[:38]
  mn := "abandon ability able about above absent absorb abstract absurd abuse access accident"
  json.NewEncoder(w).Encode(map[string]string{
   "address": addr,
   "mnemonic": mn+" "+randomHex(4),
   "privateKey": priv,
  })
 })

 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
  if r.URL.Path=="/" {
   http.ServeFile(w, r, "index.html")
   return
  }
  if r.URL.Path=="/logo.jpg" {
   http.ServeFile(w, r, "logo.jpg")
   return
  }
  http.NotFound(w,r)
 })

 log.Println("EKI Mainnet serving index.html on :"+port)
 log.Fatal(http.ListenAndServe(":"+port, nil))
}
