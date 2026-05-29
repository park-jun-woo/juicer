//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what .to(handler) 호출 인자에서 핸들러 식별자 이름을 추출한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func extractToHandler(callExpr *sitter.Node, src []byte) string {
	args := findChildByType(callExpr, "arguments")
	if args == nil {
		return ""
	}
	for i := 0; i < int(args.ChildCount()); i++ {
		if name := handlerNameFromArg(args.Child(i), src); name != "" {
			return name
		}
	}
	return ""
}
