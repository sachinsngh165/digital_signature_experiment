package main

import (
	"crypto"
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
	l, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Alice: Listening on %s", l.Addr().String())
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("connected to: %v\n", c.RemoteAddr().String())
		handleConn(c)
	}
}

func handleConn(conn net.Conn) {
	fmt.Println("Reading the message")
	encodedMessage, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	signedMessage := Message{}
	err = json.Unmarshal(encodedMessage, &signedMessage)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Reading the public key")
	publicKey, err := getPublicKey()
	if err != nil {
		fmt.Println(err)
		return
	}

	h := sha256.New()
	h.Write([]byte(signedMessage.Msg))
	hashDigest := h.Sum(nil)

	fmt.Println("Verifying the message with public key")
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashDigest, signedMessage.Sign)
	if err != nil {
		fmt.Printf("verification failed because: %+v", err)
		return
	}
	fmt.Println("Alice: Message verified and read. Love you back.")
}

func getPublicKey() (*rsa.PublicKey, error) {
	f, err := ioutil.ReadFile(".ssh/public.pem")
	if err != nil {
		return nil, err
	}

	pemBlock, rest := pem.Decode(f)
	if len(rest) > 0 {
		return nil, fmt.Errorf("unprcoessed bytes: %v", string(rest))
	}

	var iPublicKey interface{}
	if pemBlock.Type == "PUBLIC KEY" {
		iPublicKey, err = x509.ParsePKIXPublicKey(pemBlock.Bytes)
		if err != nil {
			return nil, err
		}
	}

	if iPublicKey == nil {
		return nil, errors.New("unknown pkcs found")
	}

	publicKey, ok := iPublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("failed to parse rsa private key")
	}
	return publicKey, nil
}
