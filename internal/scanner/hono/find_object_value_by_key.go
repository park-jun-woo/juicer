//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what object 리터럴에서 지정 키 pair의 value 노드를 반환한다 (없으면 nil)
package hono

import sitter "github.com/smacker/go-tree-sitter"

func findObjectValueByKey(obj *sitter.Node, key string, src []byte) *sitter.Node {
	if obj == nil {
		return nil
	}
	for _, pair := range childrenOfType(obj, "pair") {
		keyNode := pair.ChildByFieldName("key")
		if keyNode == nil || nodeText(keyNode, src) != key {
			continue
		}
		return pair.ChildByFieldName("value")
	}
	return nil
}
