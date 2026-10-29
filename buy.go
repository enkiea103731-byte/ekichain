package main
import (
 "encoding/json"
 "net/http"
)
func BalanceHandler(w http.ResponseWriter, r *http.Request) {
 wallet, err := LoadWallet()
 if err != nil {
   http.Error(w, "sem carteira", 404)
   return
 }
 json.NewEncoder(w).Encode(map[string]interface{}{
   "address": wallet.PublicAddress,
   "balance": TotalSupply,
   "price": TokenPrice,
 })
}
