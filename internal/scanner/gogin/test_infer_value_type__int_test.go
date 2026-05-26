//ff:func feature=scan type=extract control=sequence
//ff:what TestInferValueType_Int 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestInferValueType_Int(t *testing.T) {
	got := inferValueType(&ast.BasicLit{Kind: token.INT, Value: "42"}, nil)
	if got != "integer" {
		t.Fatalf("expected integer, got %s", got)
	}
}
