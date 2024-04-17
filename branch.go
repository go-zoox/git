package git

import (
	"fmt"
	"sort"

	"github.com/go-zoox/core-utils/array"
	"github.com/go-zoox/core-utils/regexp"
)

type ListBranchesConfig struct {
	WithTag bool
}

type ListBranchesOptions func(*ListBranchesConfig)

func (g *git) ListBranches(keyword string, limit int, opts ...ListBranchesOptions) ([]string, error) {
	cfg := &ListBranchesConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	refs, err := g.getRefs()
	if err != nil {
		return nil, fmt.Errorf("failed to get refs: %v", err)
	}

	branches := []string{}
	tags := []string{}
	for _, ref := range refs {
		if ref.Name().IsBranch() {
			branches = append(branches, ref.Name().String()[len("refs/heads/"):])
		}

		if cfg.WithTag {
			if ref.Name().IsTag() {
				tags = append(tags, ref.Name().String()[len("refs/tags/"):])
			}
		}
	}

	sort.Slice(branches, func(i, j int) bool {
		return branches[i] < branches[j]
	})

	// tags end
	if cfg.WithTag {
		branches = append(branches, tags...)
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
