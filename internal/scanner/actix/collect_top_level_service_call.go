//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 최상위 .service() 호출이면 그 인자에서 라우트를 수집한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func collectTopLevelServiceCall(n *sitter.Node, fi *fileInfo, routes *[]builderRoute) {
	if !isTopLevelServiceCall(n, fi.src) {
		return
	}
	args := findChildByType(n, "arguments")
	if args == nil {
		return
	}
	processServiceCallArgs(args, fi.src, "", routes)
}
