package git

import (
	"fmt"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-zoox/random"
)

func (g *git) CheckoutByBranchAndCommit(branch, commit string) (path []byte, err error) {
	w, r, err := g.getWorktree()
	if err != nil {
		return nil, fmt.Errorf("failed to get work tree: %v", err)
	}

	hashX := plumbing.NewHash(commit)

	// check and remove old branch
	_, err = g.GetHashByBranch(branch)
	if err == nil {
		// @TODO checkout to tmp branch
		tmpBranch := random.String(16)
		err = w.Checkout(&gogit.CheckoutOptions{
			Branch: plumbing.NewBranchReferenceName(tmpBranch),
			Hash:   hashX,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to checkout tmp branch: %v", err)
		}

		// remove old branch
		err = r.Storer.RemoveReference(plumbing.NewBranchReferenceName(branch))
		if err != nil {
			return nil, fmt.Errorf("failed to remove old branch: %v", err)
		}

		// remove tmp branch
		err = r.Storer.RemoveReference(plumbing.NewBranchReferenceName(tmpBranch))
		if err != nil {
			return nil, fmt.Errorf("failed to remove old branch: %v", err)
		}
	}

	// create new branch
	err = w.Checkout(&gogit.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branch),
		Hash:   hashX,
		Create: true,
		Force:  true,
		Keep:   false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to checkout new branch: %v", err)
	}

	return []byte(g.cfg.Path), nil
}

func (g *git) CheckoutByTagAndCommit(tag, commit string) (path []byte, err error) {
	w, r, err := g.getWorktree()
	if err != nil {
		return nil, fmt.Errorf("failed to get work tree: %v", err)
	}

	hashX := plumbing.NewHash(commit)

	// checkout
	hash, err := g.GetHashByBranch(tag)
	if err != nil {
		// not found branch
		err = w.Checkout(&gogit.CheckoutOptions{
			Branch: plumbing.NewTagReferenceName(tag),
			Hash:   hashX,
			Create: true,
			Force:  true,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to checkout new branch: %v", err)
		}
	} else {
		// already exist branch
		err = w.Checkout(&gogit.CheckoutOptions{
			Branch: plumbing.NewTagReferenceName(tag),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to checkout exist branch: %v", err)
		}

		if string(hash) != commit {
			ref := plumbing.NewHashReference(plumbing.NewTagReferenceName(tag), hashX)
			err = r.Storer.SetReference(ref)
			if err != nil {
				return nil, fmt.Errorf("failed to set reference: %v", err)
			}
		}
	}

	return []byte(g.cfg.Path), nil
}

// // according to action/checkout:
// //   - https://github.com/actions/checkout/blob/ac593985615ec2ede58e132d2e21d2b1cbd6127c/src/ref-helper.ts#L14
// func (g *git) getCheckoutInfo(ref, commit string) (*CheckoutInfo, error) {
// 	result := &CheckoutInfo{}
// 	if ref == "" && commit == "" {
// 		return nil, errors.New("ref and commit cannot both be empty")
// 	}

// 	// SHA only
// 	if result.Ref == "" {
// 		result.Ref = commit
// 	} else if strings.StartsWith(ref, "refs/heads/") {
// 		// refs/heads
// 		branch := ref[len("refs/heads/"):]
// 		result.Ref = branch
// 		result.StartPoint = fmt.Sprintf(`refs/remotes/origin/%s`, branch)
// 	} else if strings.StartsWith(ref, "refs/pull/") {
// 		// refs/pull/
// 		branch := ref[len("refs/pull/"):]
// 		result.Ref = fmt.Sprintf(`refs/remotes/pull/%s`, branch)
// 	} else if strings.StartsWith(ref, "refs/tags/") {
// 		// refs/tags/
// 		result.Ref = ref
// 	} else {
// 		// Unqualified ref, check for a matching branch or tag
// 		if _, err := g.GetHashByBranch(ref); err != nil {
// 			// is branch
// 			result.Ref = ref
// 			result.StartPoint = fmt.Sprintf("refs/remotes/origin/%s", ref)
// 		} else if _, err := g.GetHashByTag(ref); err != nil {
// 			// is tag
// 			result.Ref = fmt.Sprintf("refs/tags/%s", ref)
// 		} else {
// 			return nil, fmt.Errorf("A branch or tag with the name '%s' could not be found", ref)
// 		}
// 	}

// 	return result, nil
// }
