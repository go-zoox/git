package git

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/testify"
)

func TestListBranch(t *testing.T) {
	g, _ := New(&Config{
		Repository: "https://github.com/zcorky/zmicro",
		// Username:   os.Getenv("GIT_USERNAME"),
		// Password:   os.Getenv("GIT_PASSWORD"),
	})

	branches, err := g.ListBranches("", 10)
	testify.Assert(t, err == nil, "should not error")

	fmt.PrintJSON(branches)
}
