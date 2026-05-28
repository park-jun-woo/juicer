//ff:func feature=scan type=extract control=selection topic=dotnet
//ff:what 람다 파라미터를 body/query/header/path로 분류한다
package dotnet

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func classifyLambdaParam(param *sitter.Node, src []byte, path string, req *scanner.Request) {
	typeName := extractParamType(param, src)
	paramName := extractParamName(param, src)

	if isDIType(typeName) {
		return
	}

	switch {
	case hasAttribute(param, src, AttrFromBody):
		req.Body = &scanner.Body{
			VarName:  paramName,
			Method:   "FromBody",
			TypeName: typeName,
		}
	case hasAttribute(param, src, AttrFromQuery):
		req.Query = append(req.Query, scanner.Param{
			Name: paramName,
			Type: csharpTypeToOpenAPIType(typeName),
		})
	case hasAttribute(param, src, AttrFromHeader):
		req.Headers = append(req.Headers, scanner.Param{
			Name: paramName,
			Type: csharpTypeToOpenAPIType(typeName),
		})
	case isPathParam(paramName, path):
		req.PathParams = append(req.PathParams, scanner.Param{
			Name: paramName,
			Type: csharpTypeToOpenAPIType(typeName),
		})
	}
}
