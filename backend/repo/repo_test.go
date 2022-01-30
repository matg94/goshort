package repo

import (
	"fmt"
	"testing"

	"github.com/matg94/goshort/config"
)

type RepoTest struct {
	Original string
	Short    string
}

func TestRepoShorten(t *testing.T) {
	config.ReadConfig("../config/test.yaml")
	tests := []RepoTest{
		{"http://twitter.com", config.GetURLPrefix() + "3d28c09d"},
		{"http://google.com", config.GetURLPrefix() + "23498856"},
		{"http://google.com", config.GetURLPrefix() + "23498856"},
	}

	for _, test := range tests {
		shortened := ShortenURL(test.Original)
		if shortened != test.Short {
			failedLog := fmt.Sprintf("Repo shortening failed, expected %s but got %s", test.Short, shortened)
			t.Error(failedLog)
		}
	}
}

func TestRepoRetrieveOriginal(t *testing.T) {
	config.ReadConfig("../config/test.yaml")
	tests := []RepoTest{
		{"http://twitter.com", config.GetURLPrefix() + "3d28c09d"},
		{"http://google.com", config.GetURLPrefix() + "23498856"},
		{"http://google.com", config.GetURLPrefix() + "23498856"},
	}

	for _, test := range tests {
		ShortenURL(test.Original)
		original := OriginalURL(test.Short)
		if original != test.Original {
			failedLog := fmt.Sprintf("Repo shortening failed, expected %s but got %s", test.Original, original)
			t.Error(failedLog)
		}
	}
}
