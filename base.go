package git

import (
	"errors"
	"fmt"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/go-zoox/fs"
	"github.com/go-zoox/logger"
)

func (g *git) getRefs() (refs []*plumbing.Reference, err error) {
	remote := gogit.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{g.cfg.Repository},
	})

	return remote.List(&gogit.ListOptions{
		Auth:            g.auth(),
		InsecureSkipTLS: g.cfg.IsInsecureSkipTLS,
	})
}

func (g *git) auth() transport.AuthMethod {
	var auth transport.AuthMethod
	if g.cfg.Username != "" && g.cfg.Password != "" {
		auth = &http.BasicAuth{
			Username: g.cfg.Username,
			Password: g.cfg.Password,
		}
	}

	return auth
}

func (g *git) getWorktree() (w *gogit.Worktree, r *gogit.Repository, err error) {
	// open
	if fs.IsExist(g.cfg.Path) {
		r, err = gogit.PlainOpen(g.cfg.Path)
		if err != nil {
			logger.Warnf("%s not a valid git repository, removed", g.cfg.Path)
			if err = fs.RemoveDir(g.cfg.Path); err != nil {
				return nil, nil, fmt.Errorf("failed to remove invalid git repository(%s): %v", g.cfg.Path, err)
			}
		} else {
			// fetch update
			err = r.Fetch(&gogit.FetchOptions{
				Auth: g.auth(),
			})
			if err != nil {
				if !errors.Is(err, gogit.NoErrAlreadyUpToDate) {
					return nil, nil, fmt.Errorf("failed to fetch origin repository: %v", err)
				}
			}
		}
	}

	// clone
	if r == nil {
		r, err = gogit.PlainClone(g.cfg.Path, false, &gogit.CloneOptions{
			URL:  g.cfg.Repository,
			Auth: g.auth(),
		})
		if err != nil {
			return nil, nil, fmt.Errorf("failed to clone: %v", err)
		}
	}

	w, err = r.Worktree()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get work tree: %v", err)
	}

	return
}
