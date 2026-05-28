//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 단일 import 선언에서 간단한 이름과 FQCN을 추출한다
package spring

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func parseOneImport(imp *sitter.Node, src []byte) (string, string) {
	text := nodeText(imp, src)
	text = strings.TrimPrefix(text, "import ")
	text = strings.TrimSuffix(text, ";")
	text = strings.TrimSpace(text)
	if strings.HasPrefix(text, "static ") {
		return "", ""
	}
	parts := strings.Split(text, ".")
	if len(parts) == 0 {
		return "", ""
	}
	simpleName := parts[len(parts)-1]
	if simpleName == "*" {
		return "", ""
	}
	return simpleName, text
}
