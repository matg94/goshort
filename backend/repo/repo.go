package repo

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/matg94/goshort/config"
)

var fakeRedis = FakeRedis{
	store: map[string]string{},
}

func ShortenURL(originalURL string) string {
	hasher := sha1.New()
	hasher.Write([]byte(originalURL))
	sha := hex.EncodeToString(hasher.Sum(nil)[:config.GetHashLength()])

	if fakeRedis.GET(sha) == "" {
		fakeRedis.SET(sha, originalURL)
	}

	return fmt.Sprintf("%s%s", config.GetURLPrefix(), sha)
}

func OriginalURL(shortenURL string) string {
	urlArr := strings.Split(shortenURL, "/")
	sha := urlArr[len(urlArr)-1]
	originalURL := fakeRedis.GET(sha)
	return originalURL
}
