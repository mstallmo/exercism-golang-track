// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	if strings.Contains(remark, "?") {
		return "Sure."
	}
	return "Whatever."
}
