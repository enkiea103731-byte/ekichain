package main

import (
 "encoding/json"
 "net/http"
)

func handleGetSupply(w http.ResponseWriter, r *http.Request) {
 json.NewEncoder(w).Encode(map[string]any{
  "chain": "EKICHAIN",
  "symbol": "EKI",
  "total_supply": 21000000,
  "genesis": "EKI Chain Genesis Block",
  "blocks": len(Blockchain),
  "base": "BASE EKI",
 })
}
