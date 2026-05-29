//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 메서드 체인에서 .service() 호출 인자들을 콜백으로 전달한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func findServiceCalls(node *sitter.Node, src []byte, fn func(*sitter.Node)) {
	walkMethodChain(node, src, "service", fn)
}
