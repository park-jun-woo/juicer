//ff:func feature=scan type=extract control=iteration dimension=1 topic=zod
//ff:what arguments 노드에서 문자열/숫자/배열 인자를 수집한다
package zod

import sitter "github.com/smacker/go-tree-sitter"

func collectStringArgs(args *sitter.Node, src []byte) []string {
	var result []string
	argNodes := collectArgNodes(args)
	for _, a := range argNodes {
		result = append(result, extractArgValues(a, src)...)
	}
	return result
}
