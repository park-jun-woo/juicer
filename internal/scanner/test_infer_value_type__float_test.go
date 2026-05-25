//ff:func feature=scan type=extract control=sequence
//ff:what TestInferValueType_Float 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestInferValueType_Float(t *testing.T) {
	got := inferValueType(&ast.BasicLit{Kind: token.FLOAT, Value: "3.14"}, nil)
	if got != "number" {
		t.Fatalf("expected number, got %s", got)
	}
}
