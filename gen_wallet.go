package main
import (
 "crypto/ecdsa"
 "crypto/elliptic"
 "crypto/rand"
 "crypto/sha256"
 "encoding/hex"
 "fmt"
)
func main(){
 for i:=1;i<=2;i++{
   priv,_:=ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
   pubBytes:=elliptic.Marshal(elliptic.P256(), priv.X, priv.Y)
   h:=sha256.Sum256(pubBytes)
   addr:="EKI_"+hex.EncodeToString(h[:8])
   fmt.Printf("CARTEIRA %d:\nEndereço: %s\nPriv: %x\n\n",i,addr,priv.D)
 }
}
