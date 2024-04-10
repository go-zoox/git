package git

import "os"

// Git is a git client interface.
type Git interface {
	GetHashByBranch(branch string) ([]byte, error)
	GetHashByTag(tag string) ([]byte, error)
	//
	GetCommitByHash(hash string) (*Commit, error)
	//
	CheckoutByBranchAndCommit(branch, commit string) (path []byte, err error)
	CheckoutByTagAndCommit(tag, commit string) (path []byte, err error)
	//
	ListBranches(keyword string, limit int, opts ...ListBranchesOptions) ([]string, error)
}

type git struct {
	cfg *Config
}

// Config ...
type Config struct {
	Repository string `json:"repository"`
	//
	Path string `json:"path"`
	//
	Username string `json:"username"`
	Password string `json:"password"`
	//
	IsInsecureSkipTLS bool `json:"is_insecure_skip_tls"`
}

// New creates a git client.
func New(cfg *Config) (Git, error) {
	if cfg.Path == "" {
		cfg.Path = os.TempDir()
	}

	return &git{
		cfg: cfg,
	}, nil
}
