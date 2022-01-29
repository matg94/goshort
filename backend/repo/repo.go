package repo

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/matg94/goshort/config"
	"github.com/matg94/goshort/models"
)

func ShortenURL(originalURL string) models.URL {
	hasher := sha1.New()
	hasher.Write([]byte(originalURL))
	sha := hex.EncodeToString(hasher.Sum(nil)[:config.GetHashLength()])
	fmt.Println("Generated SHA:", sha)
	return models.URL{
		Original:  originalURL,
		Shortened: fmt.Sprintf("%s%s", config.GetURLPrefix(), sha),
	}
}
