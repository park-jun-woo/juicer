//ff:func feature=scan type=test control=sequence
//ff:what TestInferValueType_String 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestInferValueType_String(t *testing.T) {
	got := inferValueType(&ast.BasicLit{Kind: token.STRING, Value: `"hi"`}, nil)
	if got != "string" {
		t.Fatalf("expected string, got %s", got)
	}
}

