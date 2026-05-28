//ff:func feature=scan type=extract control=sequence topic=supafunc
//ff:what const { field1, field2 } = await req.json() 구조 분해에서 필드명을 추출한다. pair_pattern 리네이밍과 중복 제거를 지원하며, 구조분해 실패 시 dot 접근 폴백을 수행한다
package supafunc

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func extractRequestJSON(body *sitter.Node, src []byte) []string {
	seen := map[string]bool{}
	var fields []string
	walkNodes(body, func(n *sitter.Node) {
		if n.Type() != "lexical_declaration" && n.Type() != "variable_declaration" {
			return
		}
		text := nodeText(n, src)
		if !strings.Contains(text, "req.json()") {
			return
		}
		declarators := childrenOfType(n, "variable_declarator")
		for _, decl := range declarators {
			pattern := findChildByType(decl, "object_pattern")
			if pattern == nil {
				continue
			}
			for i := 0; i < int(pattern.ChildCount()); i++ {
				child := pattern.Child(i)
				var name string
				switch child.Type() {
				case "shorthand_property_identifier_pattern":
					name = nodeText(child, src)
				case "pair_pattern":
					key := findChildByType(child, "property_identifier")
					if key != nil {
						name = nodeText(key, src)
					}
				}
				if name != "" && !seen[name] {
					seen[name] = true
					fields = append(fields, name)
				}
			}
		}
	})
	if len(fields) == 0 {
		fields = extractBodyMemberAccess(body, src)
	}
	return fields
}
