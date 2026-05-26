//ff:func feature=scan type=extract control=sequence
//ff:what TestInferValueType_Nil 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestInferValueType_Nil(t *testing.T) {
	got := inferValueType(&ast.Ident{Name: "nil"}, nil)
	if got != "null" {
		t.Fatalf("expected null, got %s", got)
	}
}
