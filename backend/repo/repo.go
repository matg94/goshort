package repo

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/matg94/goshort/config"
)

func ShortenURLHash(originalURL string) string {
	hasher := sha1.New()
	hasher.Write([]byte(originalURL))
	sha := hex.EncodeToString(hasher.Sum(nil)[:config.GetHashLength()])

	if redisConn.GET(sha) == "%!s(<nil>)" {
		redisConn.SET(sha, originalURL)
	}

	return fmt.Sprintf("%s%s", config.GetURLPrefix(), sha)
}

func ShortenURLCustom(originalURL string, shortRequest string) (string, error) {
	if redisConn.GET(shortRequest) != "%!s(<nil>)" {
		return "", fmt.Errorf("URL is already in use")
	}
	redisConn.SET(shortRequest, originalURL)
	return fmt.Sprintf("%s%s", config.GetURLPrefix(), shortRequest), nil
}

func OriginalURL(shortenURL string) string {
	urlArr := strings.Split(shortenURL, "/")
	sha := urlArr[len(urlArr)-1]
	originalURL := redisConn.GET(sha)
	return originalURL
}
