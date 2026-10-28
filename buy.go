package main

import (
 "encoding/json"
 "fmt"
 "net/http"
)

type BuyRequest struct {
 From string `json:"from"`
 To string `json:"to"`
 Buyer string `json:"buyer"`
 Amount int `json:"amount"`
}

func handleBuy(w http.ResponseWriter, r *http.Request) {
 var b BuyRequest
 json.NewDecoder(r.Body).Decode(&b)
 if b.Amount == 0 { b.Amount = 1000 }
 if b.Buyer == "" { b.Buyer = b.To; if b.Buyer == "" { b.Buyer = b.From } }
 if b.Buyer == "" { b.Buyer = "EKI_CLIENTE123" }

 preco := float64(b.Amount) * 0.01
 data := fmt.Sprintf("BUY EKICHAIN | buyer:%s amount:%d preco_R$:%.2f pix:%s from:%s", b.Buyer, b.Amount, preco, b.From, b.To)
 newBlock := generateBlock(Blockchain[len(Blockchain)-1], data)
 Blockchain = append(Blockchain, newBlock)
 Balances[b.Buyer] += b.Amount

 json.NewEncoder(w).Encode(map[string]any{
  "ok": true,
  "chain": "EKICHAIN",
  "buyer": b.Buyer,
  "amount": b.Amount,
  "preco_brl": preco,
  "block": newBlock.Index,
  "hash": newBlock.Hash,
  "balance": Balances[b.Buyer],
 })
}
