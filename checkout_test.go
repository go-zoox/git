package git

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestCheckoutByBranchAndCommit(t *testing.T) {
	Path := "/tmp/go-zoox_git_test_checkout_by_hash"

	// fs.RemoveDir(Path)

	g, _ := New(&Config{
		Repository: "https://github.com/zcorky/zmicro",
		Path:       Path,
	})

	path, err := g.CheckoutByBranchAndCommit("master", "79b8bb541f072c5459ef5fe76cfeec22942abeee")
	if err != nil {
		t.Fatal(err)
	}

	testify.Equal(t, Path, string(path))
}

func TestCheckoutByTagAndCommit(t *testing.T) {
	Path := "/tmp/go-zoox_git_test_checkout_by_hash"

	// fs.RemoveDir(Path)

	g, _ := New(&Config{
		Repository: "https://github.com/zcorky/zmicro",
		Path:       Path,
	})

	path, err := g.CheckoutByTagAndCommit("v1.0.0", "e8f94341d90d53bebe736c5bc92a38950a777b5b")
	if err != nil {
		t.Fatal(err)
	}

	testify.Equal(t, Path, string(path))
}
