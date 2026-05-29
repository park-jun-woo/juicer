//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 노드의 호출 루트가 web::scope/web::resource 인지 판별한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func isWebScopeOrResource(node *sitter.Node, src []byte) bool {
	root := findCallRoot(node, src)
	return root == "web::scope" || root == "web::resource"
}
