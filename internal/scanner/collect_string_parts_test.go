//ff:func feature=scan type=test control=sequence
//ff:what TestCollectStringParts_BasicLit 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestCollectStringParts_BasicLit(t *testing.T) {
	var parts []string
	lit := &ast.BasicLit{Kind: token.STRING, Value: `"/api"`}
	collectStringParts(lit, &parts)
	if len(parts) != 1 || parts[0] != "/api" {
		t.Fatalf("got %v", parts)
	}
}
