//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 노드 하위의 모든 문자열 리터럴을 경로 세그먼트로 수집한다 (path.join 인자 대응)
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func dirStringSegments(node *sitter.Node, src []byte) []string {
	var segs []string
	for _, s := range findAllByType(node, "string") {
		segs = append(segs, unquoteTS(nodeText(s, src)))
	}
	return segs
}
