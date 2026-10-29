package main

import (
 "fmt"
 "net"
)

func StartP2P() {
 ln, err := net.Listen("tcp", ":6001")
 if err != nil {
   fmt.Println("P2P erro:", err)
   return
 }
 fmt.Println("🔗 Servidor P2P escutando na porta :6001")
 for {
   conn, _ := ln.Accept()
   conn.Close()
 }
}
