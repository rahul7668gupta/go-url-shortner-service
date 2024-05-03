package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"net/url"

	"github.com/rahul7668gupta/go-url-shortner/pkg/constants"
)

func Sha256Hash(input string) string {
	// Calculate the SHA-256 hash
	hash := sha256.Sum256([]byte(input))
	// Convert the hash to a hexadecimal string
	hashString := hex.EncodeToString(hash[:])
	return "0x" + hashString
}

func GetUrlDomain(inputUrl string) (string, error) {
	// parse url
	parsedURL, err := url.Parse(inputUrl)
	if err != nil {
		return constants.EMPTY_STRING, err
	}
	return parsedURL.Host, nil
}
