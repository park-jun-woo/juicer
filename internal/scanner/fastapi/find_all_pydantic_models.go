//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what AST에서 모든 BaseModel 서브클래스를 수집한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// findAllPydanticModels finds all BaseModel subclasses in the AST.
func findAllPydanticModels(root *sitter.Node, src []byte) map[string][]pydanticField {
	models := make(map[string][]pydanticField)
	classes := findAllByType(root, "class_definition")
	for _, cls := range classes {
		if !isBaseModelSubclass(cls, root, src) {
			continue
		}
		nameNode := findChildByType(cls, "identifier")
		if nameNode == nil {
			continue
		}
		models[nodeText(nameNode, src)] = extractPydanticFields(cls, src)
	}
	return models
}
