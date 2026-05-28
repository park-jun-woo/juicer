//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what record 타입의 생성자 파라미터를 필드로 추출한다
package dotnet

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractRecordParams(rec *sitter.Node, src []byte) []scanner.Field {
	params := findChildByType(rec, "parameter_list")
	if params == nil {
		return nil
	}
	var fields []scanner.Field
	for _, param := range childrenOfType(params, "parameter") {
		f := extractRecordParam(param, src)
		if f.Name != "" {
			fields = append(fields, f)
		}
	}
	return fields
}
