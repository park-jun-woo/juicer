//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 클래스의 public 프로퍼티를 파싱한다
package dotnet

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractClassProps(cls *sitter.Node, src []byte) []scanner.Field {
	body := findChildByType(cls, "declaration_list")
	if body == nil {
		return nil
	}
	var fields []scanner.Field
	for _, prop := range childrenOfType(body, "property_declaration") {
		f := extractOneProperty(prop, src)
		if f.Name != "" {
			fields = append(fields, f)
		}
	}
	return fields
}
