package main
import (
"encoding/json"
"fmt"
"log"
"math"
"net/http"
"os"
"strconv"
"strings"
"time"
)
type Balances map[string]float64
func loadBalances() Balances {
b := make(Balances)
data, err := os.ReadFile("balances.json")
if err != nil {
= 21000000
 b
}
json.Unmarshal(data, &b)
return b
}
func saveBalances(b Balances) {
d, _ := json.MarshalIndent(b, "", "  ")
os.WriteFile("balances.json", d, 0644)
}
func crc16(s string) string {
crc := 0xFFFF
for _, c := range []byte(s) {
^= int(c) << 8
i := 0; i < 8; i++ {
(crc & 0x8000) != 0 {
= (crc << 1) ^ 0x1021
else {
<<= 1
&= 0xFFFF
}
return fmt.Sprintf("%04X", crc)
}
func geraPixPayload(chave, nome, cidade, txid string, valor float64) string {
valorStr := fmt.Sprintf("%.2f", valor)
pix := fmt.Sprintf("0014BR.GOV.BCB.PIX01%02d%s", len(chave), chave)
payload := "000201"
payload += "26" + fmt.Sprintf("%02d", len(pix)) + pix
payload += "520400005303986"
payload += "54" + fmt.Sprintf("%02d", len(valorStr)) + valorStr
payload += "5802BR"
payload += "59" + fmt.Sprintf("%02d", len(nome)) + nome
payload += "60" + fmt.Sprintf("%02d", len(cidade)) + cidade
payload += "62" + fmt.Sprintf("%02d", len(txid)+4) + "05" + fmt.Sprintf("%02d", len(txid)) + txid
payload += "6304"
payload += crc16(payload)
return payload
}
func main() {
if _, err := os.Stat("balances.json"); os.IsNotExist(err) {
ces(loadBalances())
}
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
:= r.URL.Path
p == "/" {
r, "index.html")

strings.Contains(p, ".") {
r, "."+p)

r, "index.html")
})
http.HandleFunc("/api/balance", func(w http.ResponseWriter, r *http.Request) {
tent-Type", "application/json")
trol-Allow-Origin", "*")
:= r.URL.Query().Get("address")
addr == "" {
= "EKI_3774968e0ae21a32"
:= loadBalances()
.NewEncoder(w).Encode(map[string]interface{}{"address": addr, "balance": b[addr]})
})
http.HandleFunc("/api/pix", func(w http.ResponseWriter, r *http.Request) {
tent-Type", "application/json")
trol-Allow-Origin", "*")
td, _ := strconv.ParseFloat(r.URL.Query().Get("qtd"), 64)
qtd == 0 {
td = 1000
:= math.Round(qtd*0.01*100) / 100
 := "32468719816"
:= fmt.Sprintf("EKI%d", time.Now().Unix())
load := geraPixPayload(pixKey, "EKI CHAIN", "IBIRA", txid, valor)
rURL := "https://api.qrserver.com/v1/create-qr-code/?size=400x400&data=" + payload
.NewEncoder(w).Encode(map[string]string{"pix_key": pixKey, "valor": fmt.Sprintf("%.2f", valor), "qtd": fmt.Sprintf("%.0f", qtd), "payload": payload, "qr_url": qrURL})
})
http.HandleFunc("/api/buy", func(w http.ResponseWriter, r *http.Request) {
tent-Type", "application/json")
trol-Allow-Origin", "*")
req struct{ Address string `json:"address"`; Qtd float64 `json:"qtd"` }
.NewDecoder(r.Body).Decode(&req)
req.Address == "" { req.Address = r.URL.Query().Get("address") }
req.Qtd == 0 { req.Qtd, _ = strconv.ParseFloat(r.URL.Query().Get("qtd"), 64) }
req.Address == "" { req.Address = fmt.Sprintf("EKI_%d", time.Now().Unix()) }
:= loadBalances()
b["EKI_3774968e0ae21a32"] < req.Qtd { http.Error(w, "sem saldo", 400); return }
-= req.Qtd; b[req.Address] += req.Qtd; saveBalances(b)
.NewEncoder(w).Encode(map[string]interface{}{"status": "COMPRADO - CAIU NO ENDERECO", "address": req.Address, "qtd": req.Qtd, "balance": b[req.Address]})
})
http.HandleFunc("/api/sell", func(w http.ResponseWriter, r *http.Request) {
tent-Type", "application/json")
trol-Allow-Origin", "*")
req struct{ Address string `json:"address"`; Qtd float64 `json:"qtd"` }
.NewDecoder(r.Body).Decode(&req)
req.Address == "" { req.Address = r.URL.Query().Get("address") }
req.Qtd == 0 { req.Qtd, _ = strconv.ParseFloat(r.URL.Query().Get("qtd"), 64) }
:= loadBalances()
b[req.Address] < req.Qtd { json.NewEncoder(w).Encode(map[string]string{"status": "SALDO INSUFICIENTE"}); return }
.Address] -= req.Qtd; b["EKI_3774968e0ae21a32"] += req.Qtd; saveBalances(b)
.NewEncoder(w).Encode(map[string]interface{}{"status": "VENDIDO", "pix": req.Qtd * 0.01 * 0.95})
})
port := os.Getenv("PORT")
if port == "" { port = "8080" }
fmt.Println("RODANDO PIX 32468719816 PORTA", port)
log.Fatal(http.ListenAndServe(":"+port, nil))
}
