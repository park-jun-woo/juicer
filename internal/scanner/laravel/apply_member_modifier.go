//ff:func feature=scan type=extract control=selection topic=laravel
//ff:what member_call의 middleware()/prefix() 인자를 prefix/middleware 누적기에 적용한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func applyMemberModifier(node *sitter.Node, fi fileInfo, prefix *string, mw *[]string) {
	args := findChildByType(node, "arguments")
	if args == nil {
		return
	}
	switch lastMemberCallName(node, fi.src) {
	case "middleware":
		*mw = mergeMiddleware(*mw, extractMiddlewareArgs(args, fi))
	case "prefix":
		if p, ok := firstArgString(args, fi.src); ok {
			*prefix = joinGroupPrefix(*prefix, p)
		}
	}
}
