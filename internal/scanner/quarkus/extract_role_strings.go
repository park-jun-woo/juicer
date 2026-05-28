//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 어노테이션 인자에서 역할 문자열을 추출한다
package quarkus

import (
	"regexp"

	sitter "github.com/smacker/go-tree-sitter"
)

var roleStringRegexp = regexp.MustCompile(`['"]([^'"]+)['"]`)

func extractRoleStrings(args *sitter.Node, src []byte) []string {
	text := nodeText(args, src)
	var roles []string
	for _, m := range roleStringRegexp.FindAllStringSubmatch(text, -1) {
		if len(m) > 1 {
			roles = append(roles, m[1])
		}
	}
	return roles
}
