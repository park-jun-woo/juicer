//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what object 노드에서 지정 키의 값 노드를 찾는다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func findPairValue(obj *sitter.Node, src []byte, keyName string) *sitter.Node {
	for _, pair := range childrenOfType(obj, "pair") {
		if pairKeyName(pair, src) == keyName {
			return pairValueNode(pair)
		}
	}
	return nil
}
