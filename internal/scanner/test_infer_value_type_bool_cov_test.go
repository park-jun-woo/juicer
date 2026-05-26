//ff:func feature=scan type=test control=sequence
//ff:what TestInferValueType_BoolCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestInferValueType_BoolCov(t *testing.T) {
	if got := inferValueType(&ast.Ident{Name: "true"}, nil); got != "boolean" {
		t.Fatalf("expected boolean, got %s", got)
	}
}
