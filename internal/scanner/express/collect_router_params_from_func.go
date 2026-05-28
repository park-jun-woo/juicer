//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 단일 함수 노드의 formal_parameters에서 Router 타입 파라미터를 수집한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func collectRouterParamsFromFunc(fn *sitter.Node, src []byte, routers map[string]bool) {
	params := findChildByType(fn, "formal_parameters")
	if params == nil {
		return
	}
	for i := 0; i < int(params.NamedChildCount()); i++ {
		param := params.NamedChild(i)
		checkParamTypeAnnotation(param, src, routers)
	}
}
