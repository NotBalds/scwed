package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	os.Chdir(os.Getenv("HOME") + "/.local/share/gocwe")
	for {
		bts, _ := os.ReadFile("uuid")
		id := string(bts)
		bts, _ = os.ReadFile("key")
		key, _ := x509.ParsePKCS1PrivateKey(bts)
		data, _ := json.Marshal(User{id})
		res, _ := http.Post("http://bald.su:1337/get", "application/json", bytes.NewReader(data))
		body, _ := io.ReadAll(res.Body)
		var msgs []Message
		_ = json.Unmarshal(body, &msgs)
		var by_sender = make(map[string][]Message)
		for _, v := range msgs {
			by_sender[v.Sender] = append(by_sender[v.Sender], v)
		}
		for k := range by_sender {
			fmt.Println("From " + k + ":")
			for _, v := range by_sender[k] {
				bts, _ := base64.StdEncoding.DecodeString(v.Content)
				decr, _ := rsa.DecryptPKCS1v15(rand.Reader, key, bts)
				beeep.Notify(k, string(decr), "")
			}
		}
		time.Sleep(time.Second * 5)
	}
}
