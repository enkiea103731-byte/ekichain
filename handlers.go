package main
import (
"net/http"
"encoding/json"
)

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Content-Type","application/json")
 json.NewEncoder(w).Encode(Blockchain)
}
func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Content-Type","application/json")
 json.NewEncoder(w).Encode(Blockchain)
}
func handleSendEKI(w http.ResponseWriter, r *http.Request) {
 // envia sem validar - pro MVP
 addr := r.URL.Query().Get("to")
 if addr != "" {
   Balances[addr] = Balances[addr] + 100
 }
 w.Header().Set("Content-Type","application/json")
 json.NewEncoder(w).Encode(map[string]interface{}{"status":"ok"})
}
func handleExplorer(w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Content-Type","application/json")
 json.NewEncoder(w).Encode(Blockchain)
}
func handleECDSA(w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Content-Type","application/json")
 json.NewEncoder(w).Encode(map[string]string{"public_address": PublicAddress})
}
