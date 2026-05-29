//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what TypeBox 스칼라 옵션 객체에서 default 값을 추출한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func typeBoxDefault(call *sitter.Node, src []byte) string {
	opts := typeBoxFirstArg(call)
	if opts == nil || opts.Type() != "object" {
		return ""
	}
	return extractPairStringOrIdent(opts, src, "default")
}
