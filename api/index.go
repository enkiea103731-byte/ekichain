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

type Bal map[string]float64

func load() Bal {
	b := make(Bal)
	d, err := os.ReadFile("balances.json")
	if err != nil {
		b["EKI_3774968e0ae21a32"] = 21000000
		return b
	}
	json.Unmarshal(d, &b)
	return b
}
func save(b Bal) {
	d, _ := json.MarshalIndent(b, "", "  ")
	os.WriteFile("balances.json", d, 0644)
}
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
func main() {
	if _, err := os.Stat("balances.json"); os.IsNotExist(err) {
		save(load())
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/" {
			http.ServeFile(w, r, "index.html")
			return
		}
		if strings.Contains(path, ".") {
			http.ServeFile(w, r, "."+path)
			return
		}
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/api/balance", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		addr := r.URL.Query().Get("address")
		if addr == "" {
			addr = "EKI_3774968e0ae21a32"
		}
		b := load()
		json.NewEncoder(w).Encode(map[string]interface{}{"address": addr, "balance": b[addr]})
	})
	http.HandleFunc("/api/pix", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		qtd, _ := strconv.ParseFloat(r.URL.Query().Get("qtd"), 64)
		if qtd == 0 {
			qtd = 1000
		}
		valor := math.Round(qtd*0.01*100) / 100
		chave := "32468719816"
		txid := fmt.Sprintf("EKI%d", time.Now().Unix())
		payload := pixPayload(chave, "EKI CHAIN", "IBIRA", txid, valor)
		qr := "https://api.qrserver.com/v1/create-qr-code/?size=400x400&data=" + payload
		json.NewEncoder(w).Encode(map[string]string{"pix_key": chave, "valor": fmt.Sprintf("%.2f", valor), "qtd": fmt.Sprintf("%.0f", qtd), "payload": payload, "qr_url": qr, "txid": txid})
	})
	http.HandleFunc("/api/buy", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		var req struct {
			Address string  `json:"address"`
			Qtd     float64 `json:"qtd"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		if req.Address == "" {
			req.Address = r.URL.Query().Get("address")
		}
		if req.Qtd == 0 {
			req.Qtd, _ = strconv.ParseFloat(r.URL.Query().Get("qtd"), 64)
		}
		if req.Address == "" {
			req.Address = fmt.Sprintf("EKI_%d", time.Now().Unix())
		}
		b := load()
		if b["EKI_3774968e0ae21a32"] < req.Qtd {
			http.Error(w, "sem saldo genesis", 400)
			return
		}
		b["EKI_3774968e0ae21a32"] -= req.Qtd
		b[req.Address] += req.Qtd
		save(b)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": "COMPRADO", "address": req.Address, "qtd": req.Qtd, "balance": b[req.Address]})
	})
	http.HandleFunc("/api/sell", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		var req struct {
			Address string  `json:"address"`
			Qtd     float64 `json:"qtd"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		if req.Address == "" {
			req.Address = r.URL.Query().Get("address")
		}
		if req.Qtd == 0 {
			req.Qtd, _ = strconv.ParseFloat(r.URL.Query().Get("qtd"), 64)
		}
		b := load()
		if b[req.Address] < req.Qtd {
			json.NewEncoder(w).Encode(map[string]string{"status": "SALDO INSUFICIENTE"})
			return
		}
		b[req.Address] -= req.Qtd
		b["EKI_3774968e0ae21a32"] += req.Qtd
		save(b)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": "VENDIDO", "pix": req.Qtd * 0.01 * 0.95})
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("EKI CHAIN PIX 32468719816 PORT", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
