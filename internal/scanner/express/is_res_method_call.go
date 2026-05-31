//ff:func feature=scan type=extract control=sequence topic=express
//ff:what call_expression이 res.json/send/sendStatus/status 호출인지 판별한다
package express

import sitter "github.com/smacker/go-tree-sitter"

var resMethodNames = map[string]bool{
	"json":       true,
	"send":       true,
	"sendStatus": true,
	"render":     true,
	"redirect":   true,
}

var resObjectNames = map[string]bool{
	"res":      true,
	"response": true,
}

func isResMethodCall(call *sitter.Node, src []byte) (methodName string, ok bool) {
	if call.Type() != "call_expression" {
		return "", false
	}
	mem := findChildByType(call, "member_expression")
	if mem == nil {
		return "", false
	}
	prop := mem.ChildByFieldName("property")
	if prop == nil {
		return "", false
	}
	propName := nodeText(prop, src)

	obj := mem.ChildByFieldName("object")
	if obj == nil {
		return "", false
	}

	if obj.Type() == "identifier" && resObjectNames[nodeText(obj, src)] {
		if resMethodNames[propName] {
			return propName, true
		}
		return "", false
	}

	if obj.Type() == "call_expression" && (propName == "json" || propName == "send" || propName == "render" || propName == "redirect") {
		if isResStatusCall(obj, src) {
			return propName, true
		}
	}

	return "", false
}
