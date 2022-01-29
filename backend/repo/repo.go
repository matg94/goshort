package repo

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/matg94/goshort/config"
)

func ShortenURL(originalURL string) string {
	hasher := sha1.New()
	hasher.Write([]byte(originalURL))
	sha := hex.EncodeToString(hasher.Sum(nil)[:config.GetHashLength()])
	return fmt.Sprintf("%s%s", config.GetURLPrefix(), sha)
}

func OriginalURL(shortenURL string) string {
	return "notyetimpl.com"
}
