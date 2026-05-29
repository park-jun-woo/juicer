//ff:type feature=scan type=model topic=zod
//ff:what Zod 메서드 체인의 단일 메서드 정보 구조체
package zod

import sitter "github.com/smacker/go-tree-sitter"

// ChainMethod — Zod 메서드 체인의 단일 메서드 정보
type ChainMethod struct {
	Name string
	Args []string
	Node *sitter.Node
}
