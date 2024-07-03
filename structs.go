package main

type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"` // Base64 of ciphered content
}

type Database map[string][]Message

type User struct {
	Uuid             string `json:"uuid"` // UUID
	GetTime          string `json:"gettime"`
	GetTimeSignature string `json:"gettimesignature"`
}

type Register map[string]string

type Registration struct {
	Uuid      string `json:"uuid"`
	PublicKey string `json:"publickey"`
}

type Send struct {
	Receiver          string `json:"receiver"`
	SendTime          string `json:"sendtime"`
	SendTimeSignature string `json:"sendtimesignature"`
	Message
}
