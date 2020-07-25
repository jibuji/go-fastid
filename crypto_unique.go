package fastid

import (
	crand "crypto/rand"
	"encoding/base64"
	"log"
)

const cryptoRandSize = 24 //bytes

//CUnique unique use cryto/rand, which get the source of  /dev/urandom
func CUnique() string {
	rBytes := make([]byte, cryptoRandSize)
	n, err := crand.Read(rBytes)
	if err != nil {
		log.Printf("CUnique fall back to unique")
		return Unique(n, err)
	}
	return base64.URLEncoding.EncodeToString(rBytes)
}
