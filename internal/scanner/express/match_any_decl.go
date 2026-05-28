//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 노드가 함수/변수/export 선언이면 이름이 일치하는 본문을 반환한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func matchAnyDecl(node *sitter.Node, src []byte, name string) *sitter.Node {
	if body := matchFunctionDecl(node, src, name); body != nil {
		return body
	}
	if body := matchVariableDecl(node, src, name); body != nil {
		return body
	}
	return matchExportedDecl(node, src, name)
}
