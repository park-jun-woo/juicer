//ff:func feature=scan type=convert control=iteration dimension=1 topic=nestjs
//ff:what 경로 조각들을 결합한다
package nestjs

import "strings"

// joinParts joins path parts, ignoring empty ones.
func joinParts(parts ...string) string {
	var nonEmpty []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			nonEmpty = append(nonEmpty, p)
		}
	}
	if len(nonEmpty) == 0 {
		return "/"
	}
	result := strings.Join(nonEmpty, "/")
	for strings.Contains(result, "//") {
		result = strings.ReplaceAll(result, "//", "/")
	}
	return result
}
