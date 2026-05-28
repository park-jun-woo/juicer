//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what argument_list 노드에서 HttpStatus 또는 정수 상태 코드를 추출한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func extractStatusFromArgList(argList *sitter.Node, src []byte) string {
	for i := 0; i < int(argList.ChildCount()); i++ {
		child := argList.Child(i)
		code := matchStatusArgChild(child, src)
		if code != "" {
			return code
		}
	}
	return ""
}
