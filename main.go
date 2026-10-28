package main
import ("crypto/sha256";"encoding/hex";"fmt";"log";"net/http";"strconv";"time")
type Block struct { Index int `json:"index"`; Timestamp string `json:"timestamp"`; Data string `json:"data"`; Hash string `json:"hash"`; PrevHash string `json:"prev_hash"`; Nonce int `json:"nonce"` }
var Blockchain []Block
func calculateHash(b Block) string { r:=strconv.Itoa(b.Index)+b.Timestamp+b.Data+b.PrevHash+strconv.Itoa(b.Nonce); h:=sha256.New(); h.Write([]byte(r)); return hex.EncodeToString(h.Sum(nil)) }
func generateBlock(old Block, data string) Block { newBlock := Block{Index: old.Index+1, Timestamp: time.Now().String(), Data: data, PrevHash: old.Hash}; target := ""; for i:=0;i<difficulty;i++{ target+="0" }; for { newBlock.Hash = calculateHash(newBlock); if newBlock.Hash[:difficulty]==target { break }; newBlock.Nonce++ }; return newBlock }
func main() {
 InitWallet()
 genesisBlock := Block{Index: 0, Timestamp: time.Now().String(), Data: "EKI Genesis", PrevHash: ""}; genesisBlock.Hash = calculateHash(genesisBlock); Blockchain = append(Blockchain, genesisBlock)
 go StartP2PServer(":6001")
 http.HandleFunc("/blocks", handleGetBlockchain); http.HandleFunc("/mine", handleWriteBlock); http.HandleFunc("/balance", handleGetBalance); http.HandleFunc("/send", handleSendEKI); http.HandleFunc("/supply", handleGetSupply); http.HandleFunc("/buy", handleBuy); http.HandleFunc("/explorer", handleExplorer); http.HandleFunc("/ecdsa", handleECDSA)
 fmt.Println("\033[1;32m  🚀 EKI Chain rodando! API :8080 | P2P :6001\033[0m")
 log.Fatal(http.ListenAndServe(":8080", nil))
}
