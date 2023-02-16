package git

import "os"

type Git interface {
	GetHashByBranch(branch string) ([]byte, error)
	GetHashByTag(tag string) ([]byte, error)
	//
	GetCommitByHash(hash string) (*Commit, error)
	//
	CheckoutByBranchAndCommit(branch, commit string) (path []byte, err error)
	CheckoutByTagAndCommit(tag, commit string) (path []byte, err error)
}

type git struct {
	cfg *Config
}

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

func New(cfg *Config) (Git, error) {
	if cfg.Path == "" {
		cfg.Path = os.TempDir()
	}

	return &git{
		cfg: cfg,
	}, nil
}
