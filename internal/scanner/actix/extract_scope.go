//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what web::scope("/prefix") 스코프에서 prefix와 핸들러를 추출한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func extractScopes(fi *fileInfo) []scopeInfo {
	var scopes []scopeInfo
	walkNodes(fi.root, func(n *sitter.Node) {
		captureScope(n, fi.src, &scopes)
	})
	return scopes
}
