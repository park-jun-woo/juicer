//ff:func feature=scan type=extract control=sequence topic=express
//ff:what validate(authValidation.register) 호출의 member_expression 인자에서 크로스파일 Joi 참조를 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

// extractJoiValidatorRef — validate(<importName>.<member>) → joiValidatorRef.
// member_expression 인자가 아니면 nil을 반환한다.
func extractJoiValidatorRef(node *sitter.Node, src []byte) *joiValidatorRef {
	if node.Type() != "call_expression" {
		return nil
	}
	fn := findChildByType(node, "identifier")
	if fn == nil || !validateRequestFunctions[nodeText(fn, src)] {
		return nil
	}
	args := findChildByType(node, "arguments")
	if args == nil {
		return nil
	}
	argNodes := collectArgNodes(args)
	if len(argNodes) < 1 {
		return nil
	}
	mem := argNodes[0]
	if mem.Type() != "member_expression" {
		return nil
	}
	obj := mem.ChildByFieldName("object")
	prop := mem.ChildByFieldName("property")
	if obj == nil || prop == nil || obj.Type() != "identifier" {
		return nil
	}
	return &joiValidatorRef{ImportName: nodeText(obj, src), Member: nodeText(prop, src)}
}
