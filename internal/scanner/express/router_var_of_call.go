//ff:func feature=scan type=extract control=sequence topic=express
//ff:what router.get(...) / router.route(...) 호출에서 객체(라우터 변수)명을 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func routerVarOfCall(call *sitter.Node, src []byte) string {
	mem := findChildByType(call, "member_expression")
	if mem == nil {
		return ""
	}
	obj := findChildByType(mem, "identifier")
	if obj == nil {
		return ""
	}
	return nodeText(obj, src)
}
