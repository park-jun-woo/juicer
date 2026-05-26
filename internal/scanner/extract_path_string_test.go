//ff:func feature=scan type=test control=sequence
//ff:what TestExtractPathString_BasicLit 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestExtractPathString_BasicLit(t *testing.T) {
	expr := &ast.BasicLit{Kind: token.STRING, Value: `"/api/v1"`}
	path, ok := extractPathString(expr)
	if !ok || path != "/api/v1" {
		t.Fatalf("expected /api/v1, got %s ok=%v", path, ok)
	}
}
