//ff:func feature=scan type=extract control=selection topic=spring
//ff:what 인자 노드에서 HttpStatus 필드 접근 또는 정수 리터럴 상태 코드를 매칭한다
package spring

import (
	"strconv"

	sitter "github.com/smacker/go-tree-sitter"
)

func matchStatusArgChild(child *sitter.Node, src []byte) string {
	switch child.Type() {
	case "field_access", "identifier":
		val := nodeText(child, src)
		if code, ok := httpStatusAnnotations[val]; ok {
			return code
		}
	case "decimal_integer_literal":
		val := nodeText(child, src)
		if n, err := strconv.Atoi(val); err == nil && n >= 100 && n < 600 {
			return val
		}
	}
	return ""
}
