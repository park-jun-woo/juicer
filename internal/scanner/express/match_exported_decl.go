//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what export_statement 내부에서 이름이 일치하는 함수/변수 선언의 본문을 반환한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func matchExportedDecl(node *sitter.Node, src []byte, name string) *sitter.Node {
	if node.Type() != "export_statement" {
		return nil
	}
	for j := 0; j < int(node.ChildCount()); j++ {
		inner := node.Child(j)
		if body := matchFunctionDecl(inner, src, name); body != nil {
			return body
		}
		if body := matchVariableDecl(inner, src, name); body != nil {
			return body
		}
	}
	return nil
}
