//ff:func feature=scan type=extract control=sequence topic=supafunc
//ff:what serve() 또는 Deno.serve() 호출의 콜백 함수 본문 노드를 반환한다
package supafunc

import sitter "github.com/smacker/go-tree-sitter"

func findServeCallback(fi *fileInfo) (*sitter.Node, string) {
	var callbackBody *sitter.Node
	var handler string
	walkNodes(fi.Root, func(n *sitter.Node) {
		if callbackBody != nil {
			return
		}
		if n.Type() != "call_expression" {
			return
		}
		fnNode := findChildByType(n, "identifier")
		memberNode := findChildByType(n, "member_expression")
		if fnNode != nil && nodeText(fnNode, fi.Src) == patServe {
			handler = patServe
		} else if memberNode != nil && nodeText(memberNode, fi.Src) == patDenoServe {
			handler = patDenoServe
		} else {
			return
		}
		args := findChildByType(n, "arguments")
		if args == nil {
			return
		}
		for i := 0; i < int(args.ChildCount()); i++ {
			child := args.Child(i)
			if child.Type() == "arrow_function" || child.Type() == "function_expression" {
				body := findChildByType(child, "statement_block")
				if body != nil {
					callbackBody = body
					return
				}
			}
		}
	})
	return callbackBody, handler
}
