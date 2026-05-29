//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what 블록에서 let <varName> = TypeName{...} 바인딩의 struct 타입명을 찾는다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func resolveLetBindingType(varName string, block *sitter.Node, src []byte) string {
	for _, decl := range findAllByType(block, "let_declaration") {
		typeName := letDeclStructType(decl, varName, src)
		if typeName != "" {
			return typeName
		}
	}
	return ""
}
