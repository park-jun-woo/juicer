//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what 함수의 첫 번째 파라미터가 fastify 인스턴스 이름이면 수집한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func collectFuncParamInstance(fn *sitter.Node, src []byte, instances map[string]bool) {
	params := findChildByType(fn, "formal_parameters")
	if params == nil {
		return
	}
	name := extractFirstParamName(params, src)
	if name != "" && fastifyParamNames[name] {
		instances[name] = true
	}
}
