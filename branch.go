package git

import (
	"fmt"

	"github.com/go-zoox/core-utils/array"
	"github.com/go-zoox/core-utils/regexp"
)

func (g *git) ListBranches(keyword string, limit int) ([]string, error) {
	refs, err := g.getRefs()
	if err != nil {
		return nil, fmt.Errorf("failed to get refs: %v", err)
	}

	branches := []string{}
	for _, ref := range refs {
		if ref.Name().IsBranch() {
			if ref.Name().IsBranch() {
				branches = append(branches, ref.Name().String()[len("refs/heads/"):])
			}
		}
	}

	if keyword != "" {
		branches = array.Filter(branches, func(branch string, index int) bool {
			return regexp.Match(fmt.Sprintf("(?i)%s", keyword), branch)
		})
	}

	if limit > 0 && len(branches) > limit {
		branches = branches[:limit]
	}

	return branches, nil
}
