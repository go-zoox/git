package git

import (
	"fmt"
	"time"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

type Commit struct {
	Hash        string
	Message     string
	Author      string
	AuthorEmail string
	CreatedAt   time.Time
}

func (g *git) GetCommitByHash(hash string) (*Commit, error) {
	hashX := plumbing.NewHash(hash)

	r, err := gogit.Clone(memory.NewStorage(), nil, &gogit.CloneOptions{
		URL:  g.cfg.Repository,
		Auth: g.auth(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to clone: %v", err)
	}

	commit, err := r.CommitObject(hashX)
	return &Commit{
		Hash:        hash,
		Message:     commit.Message,
		Author:      commit.Committer.Name,
		AuthorEmail: commit.Committer.Email,
		CreatedAt:   commit.Committer.When,
	}, err
}
