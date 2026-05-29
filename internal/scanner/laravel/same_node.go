//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 두 AST 노드가 같은 소스 범위를 가리키는지 보고한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// sameNode reports whether two AST nodes refer to the same source span.
func sameNode(a, b *sitter.Node) bool {
	return a.StartByte() == b.StartByte() && a.EndByte() == b.EndByte() && a.Type() == b.Type()
}
