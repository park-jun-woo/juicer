//ff:func feature=scan type=test control=sequence
//ff:what TestCollectStringParts_SelectorExpr 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestCollectStringParts_SelectorExpr(t *testing.T) {
	// options.BaseURL + "/api/v1/admin/buildings" pattern from oapi-codegen
	var parts []string
	expr := &ast.BinaryExpr{
		Op: token.ADD,
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "options"},
			Sel: &ast.Ident{Name: "BaseURL"},
		},
		Y: &ast.BasicLit{Kind: token.STRING, Value: `"/api/v1/admin/buildings"`},
	}
	collectStringParts(expr, &parts)
	if len(parts) != 1 {
		t.Fatalf("expected 1 part, got %d: %v", len(parts), parts)
	}
	if parts[0] != "/api/v1/admin/buildings" {
		t.Fatalf("expected /api/v1/admin/buildings, got %s", parts[0])
	}
}
