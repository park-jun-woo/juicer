//ff:func feature=scan type=extract control=selection topic=laravel
//ff:what scoped/member 호출 체인을 따라 prefix와 middleware를 누적한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// walkChain walks a member_call_expression or scoped_call_expression chain
// to accumulate prefix and middleware.
func walkChain(node *sitter.Node, fi fileInfo, prefix *string, mw *[]string) {
	switch node.Type() {
	case "scoped_call_expression":
		p, m := extractGroupModifier(node, fi)
		*prefix = joinGroupPrefix(*prefix, p)
		*mw = mergeMiddleware(*mw, m)
	case "member_call_expression":
		walkMemberChain(node, fi, prefix, mw)
	}
}
