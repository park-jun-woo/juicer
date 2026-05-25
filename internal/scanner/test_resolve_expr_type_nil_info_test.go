//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveExprType_NilInfo 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestResolveExprType_NilInfo(t *testing.T) {
	typeName, fields := resolveExprType(&ast.Ident{Name: "x"}, nil)
	if typeName != "" || fields != nil {
		t.Error("expected empty for nil info")
	}
}
