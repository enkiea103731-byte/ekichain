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
type Block struct{ Index int; Timestamp string; Data string; Hash string; PrevHash string; Nonce int }
type Wallet struct{ PrivateKey string `json:"private_key"`; PublicAddress string `json:"public_address"`; Balance int64 `json:"balance"` }
var Blockchain []Block
var Balances = make(map[string]int64)
var wallet Wallet
var currentBlock = 0

func calcHash(b Block) string {
 r:=strconv.Itoa(b.Index)+b.Timestamp+b.Data+b.PrevHash+strconv.Itoa(b.Nonce)
 h:=sha256.New(); h.Write([]byte(r)); return hex.EncodeToString(h.Sum(nil))
}
func InitWallet(){
 d,err:=os.ReadFile("wallet.json")
 if err==nil { json.Unmarshal(d,&wallet); if wallet.Balance==0 { wallet.Balance=21000000 }; Balances[wallet.PublicAddress]=wallet.Balance; return }
 wallet=Wallet{PrivateKey:"eb21723ce7cd581723aeac15a7c57affe78295174351f408b7d7f686c7bd183a", PublicAddress:"EKI_e4297c5b19dbccf4dd40", Balance:21000000}
 Balances[wallet.PublicAddress]=21000000
 sd,_:=json.Marshal(wallet); os.WriteFile("wallet.json",sd,0600)
}
func StartP2P(addr string){
 l,_:=net.Listen("tcp",addr); fmt.Println("Servidor P2P escutando na porta",addr); for { c,_:=l.Accept(); if c!=nil { go func(conn net.Conn){ defer conn.Close(); b:=make([]byte,1024); conn.Read(b); conn.Write([]byte("EKI P2P OK")) }(c) } }
}

func main(){
 InitWallet()
 g:=Block{Index:0, Timestamp:time.Now().String(), Data:"EKI Genesis"}; g.Hash=calcHash(g); Blockchain=append(Blockchain,g)
 go StartP2P(":6001")
 http.HandleFunc("/balance", func(w http.ResponseWriter, r *http.Request){
  addr:=r.URL.Query().Get("address"); b:=Balances[addr]; if b==0 { b=wallet.Balance }
  json.NewEncoder(w).Encode(map[string]interface{}{"address":addr,"balance":b})
 })
 http.HandleFunc("/buy", func(w http.ResponseWriter, r *http.Request){
  var req struct{ Address string `json:"address"`; Amount int64 `json:"amount"` }
  json.NewDecoder(r.Body).Decode(&req)
  if req.Amount<=0 { req.Amount=100 }
  if req.Address=="" { req.Address=wallet.PublicAddress }
  if Balances[req.Address]==0 { Balances[req.Address]=wallet.Balance }
  if Balances[req.Address] < req.Amount { json.NewEncoder(w).Encode(map[string]interface{}{"ok":false,"balance":Balances[req.Address]}); return }
  Balances[req.Address]-=req.Amount
  wallet.Balance=Balances[req.Address]; wallet.PublicAddress=req.Address
  currentBlock++; h:=sha256.Sum256([]byte(fmt.Sprintf("%d-%s-%d",currentBlock,req.Address,req.Amount))); hash:=hex.EncodeToString(h[:])
  sd,_:=json.Marshal(wallet); os.WriteFile("wallet.json",sd,0600)
  json.NewEncoder(w).Encode(map[string]interface{}{"ok":true,"chain":"EKICHAIN","buyer":req.Address,"amount":req.Amount,"balance":Balances[req.Address],"block":currentBlock,"hash":"00"+hash[:62],"preco_brl":1})
 })
 http.HandleFunc("/blocks", func(w http.ResponseWriter, r *http.Request){ json.NewEncoder(w).Encode(Blockchain) })
 http.HandleFunc("/supply", func(w http.ResponseWriter, r *http.Request){ json.NewEncoder(w).Encode(map[string]interface{}{"supply":21000000,"balance":wallet.Balance}) })
 fmt.Println("EKI Chain API :8080 P2P :6001")
 log.Fatal(http.ListenAndServe(":8080", nil))
}
