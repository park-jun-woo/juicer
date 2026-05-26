//ff:func feature=scan type=test control=sequence
//ff:what TestExtractBinaryPath_SelectorExpr 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestExtractBinaryPath_SelectorExpr(t *testing.T) {
	// oapi-codegen pattern: options.BaseURL + "/api/v1/admin/buildings"
	e := &ast.BinaryExpr{
		Op: token.ADD,
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "options"},
			Sel: &ast.Ident{Name: "BaseURL"},
		},
		Y: &ast.BasicLit{Kind: token.STRING, Value: `"/api/v1/admin/buildings"`},
	}
	path, ok := extractBinaryPath(e)
	if !ok {
		t.Fatal("expected ok for SelectorExpr + BasicLit")
	}
	if path != "/api/v1/admin/buildings" {
		t.Fatalf("expected /api/v1/admin/buildings, got %s", path)
	}
}
