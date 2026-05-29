//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what field_declaration 노드에서 타입 노드를 찾는다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func findFieldType(fieldDecl *sitter.Node) *sitter.Node {
	for i := 0; i < int(fieldDecl.ChildCount()); i++ {
		child := fieldDecl.Child(i)
		switch child.Type() {
		case "type_identifier", "generic_type", "primitive_type",
			"scoped_type_identifier", "reference_type":
			return child
		}
	}
	return nil
}
