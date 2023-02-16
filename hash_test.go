package git

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestGetHashByBranch(t *testing.T) {
	g, _ := New(&Config{
		Repository: "https://github.com/whatwewant/whatwewant.github.io",
	})

	h, _ := g.GetHashByBranch("gh-pages")
	testify.Equal(t, "0696dc79f159e8d6b07d675d40dc88abf053244b", string(h))
}

func TestGetHashByTag(t *testing.T) {
	g, _ := New(&Config{
		Repository: "https://github.com/zcorky/zmicro",
	})

	h, _ := g.GetHashByTag("v1.0.0")
	testify.Equal(t, "e8f94341d90d53bebe736c5bc92a38950a777b5b", string(h))
}
