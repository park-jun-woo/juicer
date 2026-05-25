//ff:func feature=sql type=parse control=sequence
//ff:what "RepoName.MethodName" → "RepoMethodName" 변환 (Repository 접미사 제거)
package sqls

import (
	"strings"
)

// toQueryName converts "RepoName.MethodName" to "RepoMethodName",
// stripping "Repository" suffix from the repo name.
//
func toQueryName(id string) string {
	parts := strings.SplitN(id, ".", 2)
	if len(parts) != 2 {
		return id
	}
	repo := strings.TrimSuffix(parts[0], "Repository")
	return repo + parts[1]
}

