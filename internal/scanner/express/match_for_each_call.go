//ff:func feature=scan type=extract control=sequence topic=express
//ff:what call_expression이 <배열>.forEach(콜백) 패턴인지 확인하고 배열 변수명을 반환한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func matchForEachCall(call *sitter.Node, src []byte, routers map[string]bool) string {
	mem := findChildByType(call, "member_expression")
	if mem == nil {
		return ""
	}
	prop := mem.ChildByFieldName("property")
	if prop == nil || nodeText(prop, src) != "forEach" {
		return ""
	}
	obj := findChildByType(mem, "identifier")
	if obj == nil {
		return ""
	}
	arrVar := nodeText(obj, src)
	args := findChildByType(call, "arguments")
	if args == nil {
		return ""
	}
	fn := findChildByType(args, "arrow_function")
	if fn == nil {
		fn = findChildByType(args, "function_expression")
	}
	if fn == nil {
		return ""
	}
	if !bodyContainsRouterUse(fn, src, routers) {
		return ""
	}
	return arrVar
}
