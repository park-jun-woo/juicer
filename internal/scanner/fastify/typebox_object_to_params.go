//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what TypeBox Type.Object 인자 객체 노드를 scanner.Param 슬라이스로 변환한다 (querystring/params용)
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func typeBoxObjectToParams(objNode *sitter.Node, src []byte) []scanner.Param {
	if objNode == nil || objNode.Type() != "object" {
		return nil
	}
	var params []scanner.Param
	for _, pair := range childrenOfType(objNode, "pair") {
		p := typeBoxPairToParam(pair, src)
		if p != nil {
			params = append(params, *p)
		}
	}
	return params
}
