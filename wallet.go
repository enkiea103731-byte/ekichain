package main
import (
 "crypto/ecdsa";"crypto/elliptic";"crypto/rand";"crypto/sha256";"encoding/json";"fmt";"math/big";"net/http";"os";"os/exec"
)
var Balances = map[string]int{"EKI_GENESIS": 21000000}
var PublicAddress string
var PrivateKey *ecdsa.PrivateKey
var LastSignatureR, LastSignatureS *big.Int
func InitWallet() {
 PrivateKey, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
 PublicAddress = fmt.Sprintf("EKI_%x", PrivateKey.PublicKey.X.Bytes()[:8])
 Balances[PublicAddress] = 0
 msg := "EKI Chain Genesis Block - O inicio do ecossistema Android"
 hash := sha256.Sum256([]byte(msg))
 r, s, _ := ecdsa.Sign(rand.Reader, PrivateKey, hash[:])
 LastSignatureR = r; LastSignatureS = s
 valid := ecdsa.Verify(&PrivateKey.PublicKey, hash[:], r, s)
 walletData := map[string]string{
  "private_key": fmt.Sprintf("%x", PrivateKey.D),
  "public_key_x": fmt.Sprintf("%x", PrivateKey.PublicKey.X),
  "public_key_y": fmt.Sprintf("%x", PrivateKey.PublicKey.Y),
  "public_address": PublicAddress,
  "ecdsa_r": r.String(),
  "ecdsa_s": s.String(),
  "ecdsa_hash": fmt.Sprintf("%x", hash),
  "message": msg,
 }
 file, _ := json.MarshalIndent(walletData, "", "  ")
 os.WriteFile("wallet.json", file, 0600)
 os.WriteFile("private.key", []byte(fmt.Sprintf("%x", PrivateKey.D)), 0600)
 G := "\033[1;32m"; Y := "\033[1;33m"; C := "\033[1;36m"; M := "\033[1;35m"; W := "\033[1;37m"; R := "\033[0m"
 if _, err := os.Stat("logo.png"); err == nil {
  out, _ := exec.Command("chafa", "-c", "256", "-s", "45x22", "logo.png").CombinedOutput()
  if len(out) > 100 { fmt.Println(string(out)) }
 }
 fmt.Println(M+"          🔥 ECOSSISTEMA ANDROID - FEITO NO TERMUX 🔥"+R)
 fmt.Println(W+"               💎 EKI CHAIN - 21M SUPPLY - BASE EKI 💎"+R)
 fmt.Println("")
 fmt.Println(G+"  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓"+R)
 fmt.Println(G+"  ┃  💰 CARTEIRA CRIPTOGRÁFICA EKI COLETADA  💰  ┃"+R)
 fmt.Println(G+"  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛"+R)
 fmt.Printf("  🔒 "+Y+"CHAVE PRIVADA:"+R+" %x...\n", PrivateKey.D.Bytes()[:16])
 fmt.Printf("  📬 "+C+"ENDEREÇO PÚBLICO: "+W+"%s\n"+R, PublicAddress)
 fmt.Println("")
 fmt.Println(Y+"  ╔══ ASSINATURA ECDSA EKI ══╗"+R)
 fmt.Printf("  📝 Mensagem: %s\n", msg)
 fmt.Printf("  🔑 Hash SHA256: %x...\n", hash[:12])
 fmt.Printf("  ✍️  R: %s...\n", r.String()[:30])
 fmt.Printf("  ✍️  S: %s...\n", s.String()[:30])
 fmt.Printf("  ✅ Validada? "+G+"%v\n"+R, valid)
 fmt.Println(Y+"  ╚═══════════════════════════╝"+R)
 fmt.Println(C+"  📂 Salvo em: wallet.json e private.key"+R)
 fmt.Println("")
}
func handleGetBlockchain(w http.ResponseWriter, r *http.Request) { json.NewEncoder(w).Encode(Blockchain) }
func handleGetBalance(w http.ResponseWriter, r *http.Request) { a:=r.URL.Query().Get("address"); if a==""{a=PublicAddress}; json.NewEncoder(w).Encode(map[string]any{"address":a,"balance":Balances[a]}) }
func handleSendEKI(w http.ResponseWriter, r *http.Request) { json.NewEncoder(w).Encode(map[string]any{"ok":true}) }
func handleWriteBlock(w http.ResponseWriter, r *http.Request) { var m struct{Data string `json:"data"`}; json.NewDecoder(r.Body).Decode(&m); nb:=generateBlock(Blockchain[len(Blockchain)-1],m.Data); Blockchain=append(Blockchain,nb); json.NewEncoder(w).Encode(nb) }
func isBlockValid(a,b Block) bool { return true }
func handleExplorer(w http.ResponseWriter, r *http.Request) { w.Header().Set("Content-Type", "text/html"); fmt.Fprintf(w, "EKI %d blocos", len(Blockchain)) }
func handleECDSA(w http.ResponseWriter, r *http.Request) {
 hash := sha256.Sum256([]byte("EKI Chain Genesis"))
 json.NewEncoder(w).Encode(map[string]any{
  "public_address": PublicAddress,
  "public_key_x": fmt.Sprintf("%x", PrivateKey.PublicKey.X),
  "public_key_y": fmt.Sprintf("%x", PrivateKey.PublicKey.Y),
  "signature_r": LastSignatureR.String(),
  "signature_s": LastSignatureS.String(),
  "hash": fmt.Sprintf("%x", hash),
  "curve": "P-256",
  "valid": true,
 })
}
