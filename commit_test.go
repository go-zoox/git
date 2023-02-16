package git

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestGetCommitByHash(t *testing.T) {
	g, _ := New(&Config{
		Repository: "https://github.com/zcorky/zmicro",
	})

	c, _ := g.GetCommitByHash("e8f94341d90d53bebe736c5bc92a38950a777b5b")
	testify.Equal(t, "e8f94341d90d53bebe736c5bc92a38950a777b5b", c.Hash)
	testify.Equal(t, "chore(release): bumped version to v1.0.0\n", c.Message)
	testify.Equal(t, "Zero", c.Author)
	testify.Equal(t, "tobewhatwewant@outlook.com", c.AuthorEmail)
	testify.Equal(t, "2021-09-03 17:30:26 +0800 +0800", c.CreatedAt.String())
}
