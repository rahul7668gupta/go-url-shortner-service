package utils

import (
	"net/url"

	"github.com/rahul7668gupta/go-url-shortner/pkg/constants"
)

func GetUrlDomain(inputUrl string) (string, error) {
	// parse url
	parsedURL, err := url.Parse(inputUrl)
	if err != nil {
		return constants.EMPTY_STRING, err
	}
	return parsedURL.Host, nil
}

func GetShortCodeFromId(num int64) string {
	base62Chars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	if num == 0 {
		return string(base62Chars[0])
	}

	var result []byte

	for num > 0 {
		remainder := num % 62
		result = append([]byte{base62Chars[remainder]}, result...)
		num = num / 62
	}

	return string(result)
}
