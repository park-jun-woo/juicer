//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 반환 타입 어노테이션, response_model, status_code를 추출한다
package fastapi

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// extractReturnType extracts the return type annotation from a function definition.
// e.g., async def get_user(...) -> UserResponse:
func extractReturnType(funcDef *sitter.Node, src []byte, ri *routeInfo) {
	typeNode := findChildByType(funcDef, "type")
	if typeNode != nil {
		ri.returnType = nodeText(typeNode, src)
	}
}
