//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what AST StructType에서 필드 목록(이름/타입/json 태그)을 추출한다
package fiber

import (
	"go/ast"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func astStructFields(st *ast.StructType) []scanner.Field {
	if st.Fields == nil {
		return nil
	}
	var fields []scanner.Field
	for _, field := range st.Fields.List {
		fields = appendAstField(fields, field)
	}
	return fields
}
