package scanner

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

func TestInferValueType_Int(t *testing.T) {
	got := inferValueType(&ast.BasicLit{Kind: token.INT, Value: "42"}, nil)
	if got != "integer" {
		t.Fatalf("expected integer, got %s", got)
	}
}

func TestInferValueType_Bool(t *testing.T) {
	got := inferValueType(&ast.Ident{Name: "true"}, nil)
	if got != "boolean" {
		t.Fatalf("expected boolean, got %s", got)
	}
}

func TestInferValueType_Nil(t *testing.T) {
	got := inferValueType(&ast.Ident{Name: "nil"}, nil)
	if got != "null" {
		t.Fatalf("expected null, got %s", got)
	}
}

func TestInferValueType_Float(t *testing.T) {
	got := inferValueType(&ast.BasicLit{Kind: token.FLOAT, Value: "3.14"}, nil)
	if got != "number" {
		t.Fatalf("expected number, got %s", got)
	}
}

func TestInferValueType_SliceExprCase(t *testing.T) {
	got := inferValueType(&ast.SliceExpr{X: &ast.Ident{Name: "arr"}}, nil)
	if got != "array" {
		t.Fatalf("expected array, got %s", got)
	}
}

func TestInferValueType_Unknown(t *testing.T) {
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	got := inferValueType(&ast.Ident{Name: "x"}, info)
	if got != "unknown" {
		t.Fatalf("expected unknown, got %s", got)
	}
}
