package main

import (
 "encoding/json"
 "encoding/hex"
 "crypto/rand"
 "log"
 "net/http"
 "os"
 "github.com/tyler-smith/go-bip39"
)

func main(){
 port := os.Getenv("PORT")
 if port==""{ port="10000" }

 http.HandleFunc("/wallet/create", func(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type","application/json")
  w.Header().Set("Access-Control-Allow-Origin","*")
  entropy, _ := bip39.NewEntropy(128)
  mnemonic, _ := bip39.NewMnemonic(entropy)
  seed := bip39.NewSeed(mnemonic, "")
  priv := hex.EncodeToString(seed[:32])
  addr := "eki1"+priv[:38]
  json.NewEncoder(w).Encode(map[string]string{
   "address": addr,
   "mnemonic": mnemonic,
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
func init(){ rand.Read(make([]byte,1)) }
