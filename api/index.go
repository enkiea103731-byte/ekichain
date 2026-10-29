package handler

import (
 "encoding/json"
 "fmt"
 "math"
 "net/http"
 "os"
 "strconv"
 "strings"
 "time"
)

type Bal map[string]float64

func crc16(s string) string {
 crc := 0xFFFF
 for _, ch := range []byte(s) {
  crc ^= int(ch) << 8
  for i := 0; i < 8; i++ {
   if crc&0x8000 != 0 {
    crc = (crc << 1) ^ 0x1021
   } else {
    crc <<= 1
   }
  }
  crc &= 0xFFFF
 }
 return fmt.Sprintf("%04X", crc)
}

func pixPayload(chave, nome, cidade, txid string, valor float64) string {
 vs := fmt.Sprintf("%.2f", valor)
 pix := fmt.Sprintf("0014BR.GOV.BCB.PIX01%02d%s", len(chave), chave)
 p := "000201"
 p += "26" + fmt.Sprintf("%02d", len(pix)) + pix
 p += "520400005303986"
 p += "54" + fmt.Sprintf("%02d", len(vs)) + vs
 p += "5802BR"
 p += "59" + fmt.Sprintf("%02d", len(nome)) + nome
 p += "60" + fmt.Sprintf("%02d", len(cidade)) + cidade
 p += "62" + fmt.Sprintf("%02d", len(txid)+4) + "05" + fmt.Sprintf("%02d", len(txid)) + txid
 p += "6304"
 p += crc16(p)
 return p
}

func Handler(w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Access-Control-Allow-Origin", "*")
 path := r.URL.Path

 if strings.HasPrefix(path, "/api/pix") {
  w.Header().Set("Content-Type", "application/json")
  qtd, _ := strconv.ParseFloat(r.URL.Query().Get("qtd"), 64)
  if qtd == 0 { qtd = 1000 }
  valor := math.Round(qtd*0.01*100) / 100
  chave := "32468719816"
  txid := fmt.Sprintf("EKI%d", time.Now().Unix())
  payload := pixPayload(chave, "EKI CHAIN", "IBIRA", txid, valor)
  qr := "https://api.qrserver.com/v1/create-qr-code/?size=400x400&data=" + payload
  json.NewEncoder(w).Encode(map[string]string{"pix_key": chave, "valor": fmt.Sprintf("%.2f", valor), "qtd": fmt.Sprintf("%.0f", qtd), "payload": payload, "qr_url": qr, "txid": txid})
  return
 }

 if strings.HasPrefix(path, "/api/") {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(map[string]string{"status": "EKI CHAIN VERCEL OK", "pix": "32468719816", "supply": "21000000"})
  return
 }

 // frontend
 if _, err := os.Stat("index.html"); err == nil {
  http.ServeFile(w, r, "index.html")
  return
 }
 fmt.Fprint(w, "EKI CHAIN - https://github.com/enkiea103731-byte/ekichain")
}
