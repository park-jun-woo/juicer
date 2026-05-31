//ff:type feature=scan type=model topic=nestjs
//ff:what enum 멤버표현식 경로 해석 컨텍스트 구조체
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// enumPathCtx bundles the context needed to resolve enum member-expression
// decorator paths (same-file AST + cross-file import resolution).
type enumPathCtx struct {
	root        *sitter.Node
	src         []byte
	absFile     string
	imports     map[string]string
	projectRoot string
}
