package common_components

import (
	"math/rand"
	"time"
)

type Album struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
	Id     string  `json:"id"`
}

func GetId() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomPart := make([]byte, 10)
	for i := range randomPart {
		randomPart[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(randomPart)
}
