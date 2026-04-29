package hashing

import (
	"crypto/sha256"
	"fmt"
)

func checkHash(body string, title string) (string, string) {

	bodyHash := sha256.Sum256([]byte(body))
	titleHash := sha256.Sum256([]byte(title))

	return fmt.Sprintf("%x", bodyHash), fmt.Sprintf("%x", titleHash)
}
