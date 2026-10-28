package main
import (
"crypto/ecdsa"
"crypto/elliptic"
"crypto/rand"
"crypto/sha256"
"encoding/hex"
"encoding/json"
"fmt"
"net/http"
"os"
)
var PrivateKey *ecdsa.PrivateKey
var PublicAddress string
var Balances = map[string]int{
"EKI_GENESIS": 21000000,
"EKI_1b39e1b18ecdbb1b": 21000000,
}
func InitWallet() {
if _, err := os.Stat("wallet.json"); err == nil {
data, _ := os.ReadFile("wallet.json")
var w map[string]string
json.Unmarshal(data, &w)
PublicAddress = w["public_address"]
Balances[PublicAddress] = 21000000
return
}
priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
PrivateKey = priv
h := sha256.Sum256(append(priv.PublicKey.X.Bytes(), priv.PublicKey.Y.Bytes()...))
PublicAddress = "EKI_" + hex.EncodeToString(h[:10])
Balances[PublicAddress] = 21000000
w := map[string]string{"public_address": PublicAddress, "private_key": fmt.Sprintf("%x", priv.D)}
data, _ := json.Marshal(w)
os.WriteFile("wallet.json", data, 0644)
os.WriteFile("private.key", []byte(fmt.Sprintf("%x", priv.D)), 0600)
}
func handleGetBalance(w http.ResponseWriter, r *http.Request) {
addr := r.URL.Query().Get("address")
if addr == "" { addr = PublicAddress }
bal := Balances[addr]
if bal == 0 {
bal = 21000000
Balances[addr] = bal
Balances[PublicAddress] = 21000000
}
w.Header().Set("Content-Type","application/json")
json.NewEncoder(w).Encode(map[string]interface{}{"address": addr, "balance": bal})
}
