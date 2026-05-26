//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 단일 메서드 정의에서 엔드포인트 정보를 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractOneMethod extracts endpoint info from a single method definition.
func extractOneMethod(m *sitter.Node, src []byte, file string) (endpointInfo, bool) {
	decorators := findDecorators(m, src)
	var ep endpointInfo
	ep.file = file
	ep.line = int(m.StartPoint().Row) + 1

	foundHTTP := false
	for _, d := range decorators {
		if method, ok := httpMethods[d.name]; ok {
			ep.method = method
			ep.path = d.arg
			foundHTTP = true
		}
		if d.name == DecHttpCode {
			ep.statusCode = parseStatusCode(d.arg)
		}
		if d.name == DecUseGuards {
			ep.middleware = append(ep.middleware, d.arg)
		}
	}
	if !foundHTTP {
		return ep, false
	}

	nameNode := findChildByType(m, "property_identifier")
	if nameNode != nil {
		ep.handler = nodeText(nameNode, src)
	}

	params := extractMethodParams(m, src)
	ep.params = params.pathParams
	ep.query = params.queryParams
	ep.bodyType = params.bodyType
	ep.files = params.files
	ep.returnType = extractReturnType(m, src)

	return ep, true
}
