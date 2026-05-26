//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 같은 파일 AST에서 모델 클래스가 정의되어 있는지 확인한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// findModelInSameFile checks if a model class is defined in the same file's AST.
func findModelInSameFile(root *sitter.Node, src []byte, typeName string) bool {
	classes := findAllByType(root, "class_definition")
	for _, cls := range classes {
		nameNode := findChildByType(cls, "identifier")
		if nameNode == nil {
			continue
		}
		if nodeText(nameNode, src) == typeName {
			return true
		}
	}
	return false
}
