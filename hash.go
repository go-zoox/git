package git

import (
	"fmt"

	"github.com/go-zoox/core-utils/regexp"
)

func (g *git) GetHashByBranch(branch string) ([]byte, error) {
	re, err := regexp.New(fmt.Sprintf("^refs/heads/%s$", branch))
	// _, err := regexp.New(fmt.Sprintf("^refs/remotes/origin/%s$", branch))
	if err != nil {
		return nil, err
	}

	refs, err := g.getRefs()
	if err != nil {
		return nil, fmt.Errorf("failed to get refs: %v", err)
	}

	for _, ref := range refs {
		// fmt.Println("xxx:", ref.Name().String())
		if re.Match(ref.Name().String()) && ref.Name().IsBranch() {
			return []byte(ref.Hash().String()), nil
		}
	}

	return nil, nil
}

func (g *git) GetHashByTag(tag string) ([]byte, error) {
	re, err := regexp.New(fmt.Sprintf("^refs/tags/%s$", tag))
	if err != nil {
		return nil, err
	}

	refs, err := g.getRefs()
	if err != nil {
		return nil, fmt.Errorf("failed to get refs: %v", err)
	}

	for _, ref := range refs {
		if re.Match(ref.Name().String()) {
			return []byte(ref.Hash().String()), nil
		}
	}

	return nil, nil
}
