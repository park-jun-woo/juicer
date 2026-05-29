//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what simple_parameter의 타입 힌트(named_type 또는 primitive_type)를 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func paramTypeName(sp *sitter.Node, src []byte) string {
	typeName := ""
	if typeNode := findChildByType(sp, "named_type"); typeNode != nil {
		if nameNode := findChildByType(typeNode, "name"); nameNode != nil {
			typeName = nodeText(nameNode, src)
		}
	}
	if primType := findChildByType(sp, "primitive_type"); primType != nil {
		typeName = nodeText(primType, src)
	}
	return typeName
}
