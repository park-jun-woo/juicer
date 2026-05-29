//ff:func feature=scan type=extract control=selection topic=fastify
//ff:what autoload 옵션의 중첩 options.prefix 값을 추출한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func nestedOptionsPrefix(opts *sitter.Node, src []byte) string {
	optionsVal := findPairValue(opts, src, "options")
	switch {
	case optionsVal == nil:
		return ""
	case optionsVal.Type() != "object":
		return ""
	default:
		return extractPrefixFromOpts(optionsVal, src)
	}
}
