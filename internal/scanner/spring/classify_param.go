//ff:func feature=scan type=extract control=selection topic=spring
//ff:what 파라미터 어노테이션에 따라 path/query/body/file/form으로 분류한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func classifyParam(param *sitter.Node, src []byte, ep *endpointInfo, imports map[string]string, referrerPath, projectRoot string) {
	typeName := extractParamType(param, src)
	paramName := extractParamName(param, src)

	switch {
	case hasAnnotation(param, src, AnnPathVariable):
		classifyPathVariable(param, src, ep, typeName, paramName)
	case hasAnnotation(param, src, AnnRequestParam):
		classifyRequestParam(param, src, ep, typeName, paramName, imports, referrerPath, projectRoot)
	case hasAnnotation(param, src, AnnRequestBody):
		ep.bodyType = typeName
		ep.bodyVarName = paramName
	case hasAnnotation(param, src, AnnRequestPart) || isMultipartFile(typeName):
		classifyRequestPart(param, src, ep, paramName)
	case hasAnnotation(param, src, AnnModelAttribute):
		ep.formType = typeName
	case hasAnnotation(param, src, AnnRequestHeader):
		classifyRequestHeader(param, src, ep, typeName, paramName)
	}
}
