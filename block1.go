package main
import (
 "encoding/json"
 "fmt"
 "os"
 "time"
)
type Tx struct{
 From string `json:"from"`
 To string `json:"to"`
 Amount int `json:"amount"`
 Hash string `json:"hash"`
 Time string `json:"time"`
 Block int `json:"block"`
}
func main(){
 genesis:="EKI_3774968e0ae21a32"
 data,err:=os.ReadFile("carteiras_reais.txt")
 if err!=nil{fmt.Println("gera carteiras primeiro");return}
 // pega segunda linha que tem EKI_
 // simula tx
 tx:=Tx{
   From: genesis,
   To: "EKI_NOVA_CARTEIRA_01",
   Amount: 1000,
   Hash: "70cd416462f3aef6acfd9e47d9317cc6_block1",
   Time: time.Now().Format(time.RFC3339),
   Block: 1,
 }
 // Se tiver carteira real gerada, usa ela
 lines:=string(data)
 // salva
 f, _:=os.Create("block1.json")
 json.NewEncoder(f).Encode(tx)
 f.Close()
 f2,_:=os.Create("blocks.json")
 json.NewEncoder(f2).Encode([]Tx{tx})
 f2.Close()
 fmt.Printf("BLOCK #1 CRIADO: %s -> %s : 1000 EKI\nHash: %s\n",tx.From,tx.To,tx.Hash)
}
