//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what autoload 옵션의 베이스 prefix(options.prefix 또는 prefix)를 추출한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func extractAutoloadPrefix(opts *sitter.Node, src []byte) string {
	if p := nestedOptionsPrefix(opts, src); p != "" {
		return p
	}
	return extractPrefixFromOpts(opts, src)
}
