package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

type ecdsaGen struct {
	curve elliptic.Curve
}

func (e *ecdsaGen) KeyGen() (key *ecdsa.PrivateKey, err error) {
	privKey, err := ecdsa.GenerateKey(e.curve, rand.Reader)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	// 生成ecdsa
	e := &ecdsaGen{curve: elliptic.P256()}
	priKey, err := e.KeyGen()
	checkError(err)

	priKeyEncode, err := x509.MarshalECPrivateKey(priKey)
	checkError(err)
	f, err := os.Create("ec.pem")
	checkError(err)
	pem.Encode(f, &pem.Block{Type: "EC PRIVATE KEY", Bytes: priKeyEncode})
	f.Close()
}
