//ff:func feature=scan type=test control=sequence
//ff:what TestCollectStringParts_Ident 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestCollectStringParts_Ident(t *testing.T) {
	// baseURL + "/path" pattern — variable reference on left side
	var parts []string
	expr := &ast.BinaryExpr{
		Op: token.ADD,
		X:  &ast.Ident{Name: "baseURL"},
		Y:  &ast.BasicLit{Kind: token.STRING, Value: `"/api/v1/users"`},
	}
	collectStringParts(expr, &parts)
	if len(parts) != 1 {
		t.Fatalf("expected 1 part, got %d: %v", len(parts), parts)
	}
	if parts[0] != "/api/v1/users" {
		t.Fatalf("expected /api/v1/users, got %s", parts[0])
	}
}
