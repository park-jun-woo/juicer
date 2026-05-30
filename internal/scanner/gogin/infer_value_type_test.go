//ff:func feature=scan type=test control=sequence
//ff:what TestInferValueType_String 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestInferValueType_String(t *testing.T) {
	got := inferValueType(&ast.BasicLit{Kind: token.STRING, Value: `"hi"`}, nil)
	if got != "string" {
		t.Fatalf("expected string, got %s", got)
	}
}

func TestInferValueType_NonGinComposite(t *testing.T) {
	// a slice composite literal -> array (isGinH false path then array detection)
	lit := &ast.CompositeLit{Type: &ast.ArrayType{Elt: &ast.Ident{Name: "int"}}}
	info := &types.Info{Types: map[ast.Expr]types.TypeAndValue{}}
	if got := inferValueType(lit, info); got != "array" {
		t.Fatalf("expected array, got %q", got)
	}
}
