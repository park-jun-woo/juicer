//ff:func feature=scan type=test control=sequence
//ff:what paramFieldAtIndex 전 분기 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestParamFieldAtIndex(t *testing.T) {
	params := &ast.FieldList{
		List: []*ast.Field{
			{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "int"}},
			{Names: []*ast.Ident{{Name: "b"}, {Name: "c"}}, Type: &ast.Ident{Name: "string"}},
			{Type: &ast.Ident{Name: "float64"}}, // unnamed
		},
	}

	f, n := paramFieldAtIndex(params, 0)
	if f == nil || n != "a" {
		t.Fatalf("idx 0: got %v, %q", f, n)
	}

	f, n = paramFieldAtIndex(params, 1)
	if f == nil || n != "b" {
		t.Fatalf("idx 1: got %v, %q", f, n)
	}

	f, n = paramFieldAtIndex(params, 2)
	if f == nil || n != "c" {
		t.Fatalf("idx 2: got %v, %q", f, n)
	}

	f, n = paramFieldAtIndex(params, 3)
	if f == nil || n != "_" {
		t.Fatalf("idx 3: got %v, %q", f, n)
	}

	// out of range
	f, n = paramFieldAtIndex(params, 99)
	if f != nil || n != "" {
		t.Fatal("expected nil for out of range")
	}
}
