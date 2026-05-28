//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what 옵션 객체에서 prefix 문자열을 추출한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func extractPrefixFromOpts(obj *sitter.Node, src []byte) string {
	val := findPairValue(obj, src, "prefix")
	if val == nil {
		return ""
	}
	if val.Type() == "string" || val.Type() == "template_string" {
		return unquoteTS(nodeText(val, src))
	}
	return ""
}
