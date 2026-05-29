//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what .route() 인자 노드에서 (method, handler)를 파싱한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func parseRouteArg(args *sitter.Node, src []byte) (string, string) {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		fe := routeToFieldExpr(child, src)
		if fe == nil {
			continue
		}
		return extractWebMethod(fe, src), extractToHandler(child, src)
	}
	return "", ""
}
