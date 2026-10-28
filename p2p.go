package main
import ("fmt";"net")
const difficulty = 2
func BroadcastBlock(b Block) {}
func StartP2PServer(port string) {
  ln, err := net.Listen("tcp", port)
  if err != nil {
    fmt.Printf("⚠️  P2P %s já em uso, pulando...\n", port)
    return
  }
  fmt.Println("🔗 Servidor P2P escutando na porta", port)
  for { conn, _ := ln.Accept(); conn.Close() }
}
