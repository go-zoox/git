package git

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestGetHashByBranch(t *testing.T) {
	g, _ := New(&Config{
		Repository: "https://github.com/whatwewant/whatwewant.github.io",
	})

	h, err := g.GetHashByBranch("gh-pages")
	if err != nil {
		t.Fatal(err)
	}

	testify.Equal(t, "0696dc79f159e8d6b07d675d40dc88abf053244b", string(h))

	_, err = g.GetHashByBranch("not-found-branch")
	testify.Assert(t, err != nil, "branch not-found-branch should not found")
	testify.Equal(t, err.Error(), "branch not-found-branch not found")
}

func TestGetHashByTag(t *testing.T) {
	g, _ := New(&Config{
		Repository: "https://github.com/zcorky/zmicro",
	})

	h, _ := g.GetHashByTag("v1.0.0")
	testify.Equal(t, "e8f94341d90d53bebe736c5bc92a38950a777b5b", string(h))
}
