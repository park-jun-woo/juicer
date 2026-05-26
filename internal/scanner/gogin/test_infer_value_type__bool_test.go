//ff:func feature=scan type=extract control=sequence
//ff:what TestInferValueType_Bool 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestInferValueType_Bool(t *testing.T) {
	got := inferValueType(&ast.Ident{Name: "true"}, nil)
	if got != "boolean" {
		t.Fatalf("expected boolean, got %s", got)
	}
}
