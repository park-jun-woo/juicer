//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what member_call 한 단계의 prefix/middleware를 적용하고 안쪽 체인으로 재귀한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func walkMemberChain(node *sitter.Node, fi fileInfo, prefix *string, mw *[]string) {
	applyMemberModifier(node, fi, prefix, mw)
	if inner := findChildByType(node, "scoped_call_expression"); inner != nil {
		walkChain(inner, fi, prefix, mw)
	}
	if innerMC := findChildByType(node, "member_call_expression"); innerMC != nil {
		walkChain(innerMC, fi, prefix, mw)
	}
}
