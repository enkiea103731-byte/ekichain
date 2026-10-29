package main

import (
 "encoding/json"
 "fmt"
 "log"
 "net/http"
 "os"
 "time"
)

type ChainInfo struct {
 Name string `json:"name"`
 ChainID int `json:"chainId"`
 Pix string `json:"pix"`
 Status string `json:"status"`
 BlockHeight int `json:"blockHeight"`
 RPC string `json:"rpc"`
 Telegram string `json:"telegram"`
}

func enableCORS(w http.ResponseWriter){
 w.Header().Set("Access-Control-Allow-Origin","*")
 w.Header().Set("Access-Control-Allow-Methods","GET, POST, OPTIONS")
 w.Header().Set("Access-Control-Allow-Headers","Content-Type")
 w.Header().Set("Content-Type","application/json")
}

func main(){
 fmt.Println("=================================")
 fmt.Println("  EKI CHAIN - NODE OFICIAL")
 fmt.Println("  PIX: 32468719816")
 fmt.Println("  ChainID: 701 (0x2BD)")
 fmt.Println("  100% BRASILEIRA - IBIRA SP")
 fmt.Println("  1a Blockchain Android do Mundo")
 fmt.Println("=================================")

 // Carrega balances
 balances := map[string]interface{}{}
 if data, err := os.ReadFile("balances.json"); err == nil {
  json.Unmarshal(data, &balances)
  fmt.Printf("Balances carregado: %d carteiras\n", len(balances))
 }

 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
  enableCORS(w)
  if r.Method == "OPTIONS" { return }
  info := ChainInfo{
   Name: "EKIChain Mainnet",
   ChainID: 701,
   Pix: "32468719816",
   Status: "Mainnet Live - Android Native",
   BlockHeight: int(time.Now().Unix() / 10),
   RPC: "https://eki-chain.vercel.app/api/rpc",
   Telegram: "https://t.me/+_RKJG6QGhqUzZTBh",
  }
  json.NewEncoder(w).Encode(info)
 })

 http.HandleFunc("/api/rpc", func(w http.ResponseWriter, r *http.Request){
  enableCORS(w)
  json.NewEncoder(w).Encode(map[string]interface{}{
   "jsonrpc":"2.0","id":1,
   "result": map[string]interface{}{
    "chainId":"0x2BD","chainName":"EKIChain",
    "nativeCurrency": map[string]interface{}{"name":"EKI","symbol":"EKI","decimals":2},
   },
  })
 })

 http.HandleFunc("/balances", func(w http.ResponseWriter, r *http.Request){
  enableCORS(w)
  if data, err := os.ReadFile("balances.json"); err == nil {
   w.Write(data)
  } else {
   json.NewEncoder(w).Encode(balances)
  }
 })

 port := "8080"
 if p := os.Getenv("PORT"); p != "" { port = p }
 fmt.Printf("EKI CHAIN NODE RODANDO NA PORTA %s\n", port)
 fmt.Printf("RPC: http://localhost:%s/api/rpc\n", port)
 fmt.Printf("Explorer: http://localhost:%s/balances\n", port)
 fmt.Printf("Telegram: https://t.me/+_RKJG6QGhqUzZTBh\n")
 log.Fatal(http.ListenAndServe(":"+port, nil))
}
