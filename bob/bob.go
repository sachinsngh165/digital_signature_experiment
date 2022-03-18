package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
)

type Message struct {
	Msg  string `json:"message"`
	Sign []byte `json:"sign"`
}

func main() {
	fmt.Println("Creating message")
	message := "Hello, Alice"
	h := sha256.New()
	h.Write([]byte(message))
	hashDigest := h.Sum(nil)

	fmt.Println("Reading private key")
	privateKey, err := getPrivateKey()
	if err != nil {
		fmt.Printf("failed to fetch private key. error: %v", err)
		return
	}

	fmt.Println("Signing the message with private key")
	messageSign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashDigest)
	if err != nil {
		fmt.Println(err)
		return
	}

	signedMessage := Message{
		Msg:  message,
		Sign: messageSign,
	}

	encodedMessage, err := json.Marshal(signedMessage)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Bob: Sending the signed message over localhost:8081")
	c, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	_, err = c.Write(encodedMessage)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Bob: Message sent")
}

func getPrivateKey() (*rsa.PrivateKey, error) {
	f, err := ioutil.ReadFile(".ssh/private.pem")
	if err != nil {
		return nil, err
	}

	pemBlock, rest := pem.Decode(f)
	if len(rest) > 0 {
		return nil, fmt.Errorf("unprcoessed bytes: %v", string(rest))
	}

	var iPrivateKey interface{}
	if pemBlock.Type == "PRIVATE KEY" {
		iPrivateKey, err = x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
		if err != nil {
			return nil, err
		}
	}

	if iPrivateKey == nil {
		return nil, errors.New("unknown pkcs found")
	}

	privateKey, ok := iPrivateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("failed to parse rsa private key")
	}
	return privateKey, nil
}
