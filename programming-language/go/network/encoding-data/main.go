package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type userData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {
	user := userData{
		Name:     "guest1",
		Password: "guest1234",
	}

	hash := sha256.Sum256([]byte(user.Password))

	// Encode Result : 674c4c7106319650affbf1b8e75a507acc405bbaeb4d56ae982ab0a68b2cce8c
	encoded := hex.EncodeToString(hash[:])
	fmt.Println(encoded)

	// Decode Result : [103 76 76 113 6 49 150 80 175 191 241 184 231 90 80 122 204 64 91 186 235 77 86 174 152 42 176 166 139 44 206 140]
	decoded, _ := hex.DecodeString(encoded)
	fmt.Println(decoded)
}
