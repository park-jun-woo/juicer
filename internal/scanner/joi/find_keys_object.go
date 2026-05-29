//ff:func feature=scan type=extract control=sequence topic=joi
//ff:what Joi.object().keys({...}) / Joi.object({...}) 에서 필드 정의 object 리터럴을 찾는다
package joi

import sitter "github.com/smacker/go-tree-sitter"

// joiSchemaMethods — 필드 object를 인자로 받는 Joi 메서드명
var joiSchemaMethods = map[string]bool{
	"keys":   true,
	"object": true,
	"append": true,
}

// FindKeysObject — Joi.object().keys({...}) 체인에서 필드 object 노드를 찾는다.
// keys/object/append 메서드 호출의 첫 object 인자를 반환한다 (없으면 nil).
func FindKeysObject(node *sitter.Node, src []byte) *sitter.Node {
	var found *sitter.Node
	walkNodes(node, func(n *sitter.Node) {
		if found != nil || n.Type() != "call_expression" {
			return
		}
		fn := n.ChildByFieldName("function")
		if fn == nil || fn.Type() != "member_expression" {
			return
		}
		prop := fn.ChildByFieldName("property")
		if prop == nil || !joiSchemaMethods[nodeText(prop, src)] {
			return
		}
		args := findChildByType(n, "arguments")
		if args == nil {
			return
		}
		if obj := findChildByType(args, "object"); obj != nil {
			found = obj
		}
	})
	return found
}
