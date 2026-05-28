//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what AST에서 클래스 선언을 찾아 필드와 타입 파라미터를 해석한다
package quarkus

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func resolveClassFromASTWithParams(root *sitter.Node, src []byte, className, filePath, projectRoot string, cache map[string][]scanner.Field) ([]scanner.Field, []string) {
	imports := extractImports(root, src)
	classes := findAllByType(root, "class_declaration")
	for _, cls := range classes {
		nameNode := findChildByType(cls, "identifier")
		if nameNode == nil || nodeText(nameNode, src) != className {
			continue
		}
		typeParams := extractTypeParams(cls, src)
		fields := extractClassFields(cls, src)
		parentFields := resolveParentFields(cls, src, filePath, projectRoot, imports, cache)
		combined := mergeParentFields(parentFields, fields)
		combined = convertFieldTypes(combined)
		cache[className] = combined
		return combined, typeParams
	}
	return nil, nil
}
