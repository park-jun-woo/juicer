//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what record 파라미터에서 필드 정보를 추출한다
package dotnet

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractRecordParam(param *sitter.Node, src []byte) scanner.Field {
	typeName := extractParamType(param, src)
	paramName := extractParamName(param, src)

	f := scanner.Field{
		Name: paramName,
		JSON: paramName,
	}
	if typeName != "" {
		raw, nullable := stripNullable(typeName)
		oa := csharpTypeToOpenAPI(raw)
		f.Type = oa.Type
		f.Nullable = nullable
	}
	applyDataAnnotations(param, src, &f)
	return f
}
