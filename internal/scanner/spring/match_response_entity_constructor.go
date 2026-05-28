//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what new ResponseEntity 생성자 노드에서 HttpStatus 상태 코드를 추출한다
package spring

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func matchResponseEntityConstructor(obj *sitter.Node, src []byte) string {
	text := nodeText(obj, src)
	if !strings.Contains(text, "ResponseEntity") {
		return ""
	}
	args := findChildByType(obj, "argument_list")
	if args == nil {
		return ""
	}
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() != "field_access" && child.Type() != "identifier" {
			continue
		}
		val := nodeText(child, src)
		if code, ok := httpStatusAnnotations[val]; ok {
			return code
		}
	}
	return ""
}
