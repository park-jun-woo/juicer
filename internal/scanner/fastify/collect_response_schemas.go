//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what response 객체에서 상태 코드별 스키마를 수집한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func collectResponseSchemas(si *schemaInfo, val *sitter.Node, src []byte) {
	if val.Type() != "object" {
		return
	}
	for _, rpair := range childrenOfType(val, "pair") {
		statusCode := pairKeyName(rpair, src)
		rval := pairValueNode(rpair)
		if statusCode != "" && rval != nil {
			si.Response[statusCode] = rval
		}
	}
}
