package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"

	"github.com/nohj0518/hyeonjucoin-2021/utils"
)

func Start() {
	privateKey,err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)

	message:= "i love you"
	hashedMessage := utils.Hash(message)

	hashAsBytes, err :=hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	r, s, err:= ecdsa.Sign(rand.Reader, privateKey,hashAsBytes)
	utils.HandleErr(err)

}










/*
1) we hash the msg.
	"i love you" -> hash(msg)
	-> "hashed_message"
2) generate key pair
	KeyPair(privateKey, publicKey)
	-> save privateKey to a file == Wallet
3) sign the hash
	("hashed_message"+privateKey)
	-> "signature"
4) verify
	("hashed_message" + "signature" + publicKey)
	-> true
*/