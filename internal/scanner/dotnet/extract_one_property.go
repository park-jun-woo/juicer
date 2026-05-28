//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 프로퍼티 노드에서 필드 정보를 추출한다
package dotnet

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractOneProperty(prop *sitter.Node, src []byte) scanner.Field {
	var f scanner.Field
	typeName := extractPropertyTypeName(prop, src)
	f.Name = extractPropertyName(prop, src)
	f.JSON = f.Name

	if typeName != "" {
		raw, nullable := stripNullable(typeName)
		oa := csharpTypeToOpenAPI(raw)
		f.Type = oa.Type
		f.Nullable = nullable
	}

	applyDataAnnotations(prop, src, &f)
	return f
}
