package main

import (
 "crypto/ecdsa"
 "crypto/elliptic"
 "crypto/rand"
 "crypto/sha256"
 "encoding/hex"
 "encoding/json"
 "math/big"
 "os"
)

type Wallet struct {
 PrivateKey string `json:"private_key"`
 PublicAddress string `json:"public_address"`
 PublicKey string `json:"public_key"`
}

func NewWallet() (*Wallet, *ecdsa.PrivateKey, error) {
 priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
 if err != nil { return nil, nil, err }
 d := priv.D.Bytes()
 privHex := hex.EncodeToString(d)
 pubBytes := elliptic.Marshal(elliptic.P256(), priv.PublicKey.X, priv.PublicKey.Y)
 pubHex := hex.EncodeToString(pubBytes)
 h := sha256.Sum256([]byte(pubHex))
 addr := "EKI_" + hex.EncodeToString(h[:])[:16]
 w := &Wallet{PrivateKey: privHex, PublicAddress: addr, PublicKey: pubHex}
 j, _ := json.MarshalIndent(w, "", "  ")
 os.WriteFile("wallet.json", j, 0644)
 os.WriteFile("private.key", []byte(privHex), 0600)
 return w, priv, nil
}

func LoadWallet() (*Wallet, error) {
 b, err := os.ReadFile("wallet.json")
 if err != nil { return nil, err }
 var w Wallet
 json.Unmarshal(b, &w)
 return &w, nil
}

func SignAndVerify(priv *ecdsa.PrivateKey, msg string) (string, string, bool) {
 hash := sha256.Sum256([]byte(msg))
 r, s, _ := ecdsa.Sign(rand.Reader, priv, hash[:])
 ok := ecdsa.Verify(&priv.PublicKey, hash[:], r, s)
 return hex.EncodeToString(r.Bytes()), hex.EncodeToString(s.Bytes()), ok
}

func GetPrivFromHex(hexStr string) *ecdsa.PrivateKey {
 b, _ := hex.DecodeString(hexStr)
 d := new(big.Int).SetBytes(b)
 priv := new(ecdsa.PrivateKey)
 priv.PublicKey.Curve = elliptic.P256()
 priv.D = d
 priv.PublicKey.X, priv.PublicKey.Y = elliptic.P256().ScalarBaseMult(d.Bytes())
 return priv
}
