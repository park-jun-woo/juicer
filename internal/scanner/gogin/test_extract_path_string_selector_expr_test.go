//ff:func feature=scan type=test control=sequence
//ff:what TestExtractPathString_SelectorExpr 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestExtractPathString_SelectorExpr(t *testing.T) {
	// oapi-codegen pattern: options.BaseURL + "/api/v1/admin/buildings"
	bin := &ast.BinaryExpr{
		Op: token.ADD,
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "options"},
			Sel: &ast.Ident{Name: "BaseURL"},
		},
		Y: &ast.BasicLit{Kind: token.STRING, Value: `"/api/v1/admin/buildings"`},
	}
	path, ok := extractPathString(bin)
	if !ok {
		t.Fatal("expected ok for SelectorExpr + BasicLit binary expr")
	}
	if path != "/api/v1/admin/buildings" {
		t.Fatalf("expected /api/v1/admin/buildings, got %s", path)
	}
}
