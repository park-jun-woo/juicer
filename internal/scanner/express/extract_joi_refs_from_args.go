//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 미들웨어 인자(argNodes[1..n-1])에서 validate(x.y) 크로스파일 Joi 참조를 찾는다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractJoiRefsFromArgs(argNodes []*sitter.Node, src []byte) []joiValidatorRef {
	if len(argNodes) < 2 {
		return nil
	}
	var refs []joiValidatorRef
	for i := 1; i < len(argNodes); i++ {
		node := argNodes[i]
		if node.Type() != "call_expression" {
			continue
		}
		if ref := extractJoiValidatorRef(node, src); ref != nil {
			refs = append(refs, *ref)
		}
	}
	return refs
}
