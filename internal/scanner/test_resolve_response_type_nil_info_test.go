//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveResponseType_NilInfo 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestResolveResponseType_NilInfo(t *testing.T) {
	typeName, fields, confidence := resolveResponseType(&ast.Ident{Name: "x"}, nil)
	if typeName != "" || fields != nil || confidence != "" {
		t.Error("expected empty for nil info")
	}
}
