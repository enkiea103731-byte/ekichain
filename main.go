cd ~/eki-chain
cat > main.go << 'EOF'
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Block struct {
	Index     int    `json:"index"`
	Timestamp string `json:"timestamp"`
	Data      string `json:"data"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prev_hash"`
	Nonce     int    `json:"nonce"`
}

type Wallet struct {
	PrivateKey    string `json:"private_key"`
	PublicAddress string `json:"public_address"`
	Balance       int64  `json:"balance"`
}

var Blockchain []Block
var Balances = map[string]int64{}
var wallet Wallet
var currentBlock = 0

func calcHash(b Block) string {
	r := strconv.Itoa(b.Index) + b.Data + b.PrevHash + strconv.Itoa(b.Nonce)
	h := sha256.New()
	h.Write([]byte(r))
	return hex.EncodeToString(h.Sum(nil))
}

func InitWallet() {
	d, err := os.ReadFile("wallet.json")
	if err == nil {
		json.Unmarshal(d, &wallet)
		if wallet.Balance == 0 {
			wallet.Balance = 21000000
		}
		Balances[wallet.PublicAddress] = wallet.Balance
		return
	}
	// genesis fixo
	wallet = Wallet{
		PrivateKey:    "a61272ec7f681723aecc15a27a4ff78295745140bb47f66e7d9d81a",
		PublicAddress: "EKI_e4297c5b19dbccf4dd40",
		Balance:       21000000,
	}
	Balances[wallet.PublicAddress] = 21000000
	sd, _ := json.Marshal(wallet)
	os.WriteFile("wallet.json", sd, 0600)
}

func StartP2P(addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("P2P erro:", err)
		return
	}
	fmt.Println("Servidor P2P escutando na porta", addr)
	for {
		c, _ := l.Accept()
		if c != nil {
			go func(conn net.Conn) {
				defer conn.Close()
				conn.Write([]byte("EKI P2P OK\n"))
			}(c)
		}
	}
}

func main() {
	InitWallet()
	g := Block{Index: 0, Timestamp: time.Now().String(), Data: "EKI Genesis"}
	g.Hash = calcHash(g)
	Blockchain = append(Blockchain, g)
	go StartP2P(":6001")

	http.HandleFunc("/balance", func(w http.ResponseWriter, r *http.Request) {
		addr := r.URL.Query().Get("address")
		b := Balances[addr]
		if b == 0 && addr == wallet.PublicAddress {
			b = wallet.Balance
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"address": addr, "balance": b})
	})

	http.HandleFunc("/buy", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Address string `json:"address"`
			Amount  int64  `json:"amount"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		if req.Amount == 0 {
			req.Amount = 100
		}
		if req.Address == "" {
			req.Address = wallet.PublicAddress
		}
		if Balances[req.Address] == 0 {
			Balances[req.Address] = Balances[wallet.PublicAddress]
		}
		if Balances[req.Address] < req.Amount {
			json.NewEncoder(w).Encode(map[string]interface{}{"ok": false, "balance": Balances[req.Address]})
			return
		}
		Balances[req.Address] -= req.Amount
		if req.Address == wallet.PublicAddress {
			wallet.Balance = Balances[req.Address]
			sd, _ := json.Marshal(wallet)
			os.WriteFile("wallet.json", sd, 0600)
		}
		currentBlock++
		bk := Block{Index: currentBlock, Timestamp: time.Now().String(), Data: fmt.Sprintf("buy %s %d", req.Address, req.Amount), PrevHash: Blockchain[len(Blockchain)-1].Hash}
		bk.Hash = calcHash(bk)
		Blockchain = append(Blockchain, bk)
		h := sha256.Sum256([]byte(fmt.Sprintf("%d%s%d", currentBlock, req.Address, req.Amount)))
		hash := hex.EncodeToString(h[:])
		json.NewEncoder(w).Encode(map[string]interface{}{
			"ok": true, "chain": "EKICHAIN", "buyer": req.Address, "amount": req.Amount,
			"balance": Balances[req.Address], "block": currentBlock, "hash": hash,
			"preco_brl": 1,
		})
	})

	http.HandleFunc("/blocks", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Blockchain)
	})

	http.HandleFunc("/supply", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"supply": 21000000, "balance": Balances[wallet.PublicAddress]})
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("EKI Chain API online - /supply /balance /buy /blocks"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("EKI Chain API 0.0.0.0:" + port + " P2P :6001")
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
EOF
cat main.go
