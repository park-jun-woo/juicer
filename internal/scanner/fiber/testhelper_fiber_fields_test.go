//ff:func feature=scan type=test topic=fiber control=sequence
//ff:what fiber 테스트 헬퍼 — Go 소스 파싱 후 첫 struct의 필드 리스트 반환
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func fiberFields(t *testing.T, src string) []*ast.Field {
	t.Helper()
	f, err := parser.ParseFile(token.NewFileSet(), "x.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	var fields []*ast.Field
	ast.Inspect(f, func(n ast.Node) bool {
		if st, ok := n.(*ast.StructType); ok {
			fields = st.Fields.List
			return false
		}
		return true
	})
	return fields
}
