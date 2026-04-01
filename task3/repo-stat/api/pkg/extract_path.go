package pkg

import (
	"fmt"
	"strings"
)

func ExtractPath(link string) (string, string, error) {
	endpoints := strings.Split(link, "/")
	var owner, repo string
	var err error

	switch len(endpoints) {
	case 2: // owner/repo
		owner, repo = endpoints[0], endpoints[1]
	case 5: // https://github.com/owner/repo
		owner, repo = endpoints[3], endpoints[4]
	default:
		err = fmt.Errorf("extract path: incorrect link")
	}

	return owner, repo, err
}
