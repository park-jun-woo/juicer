//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestResolveClassFromAST_Round5 테스트
package spring

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveClassFromAST_Round5(t *testing.T) {
	root, src := sParse(t, `class UserDto { public String name; public int age; }`)
	fields := resolveClassFromAST(root, src, "UserDto", "C.java", "/abs", map[string][]scanner.Field{})
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d: %+v", len(fields), fields)
	}
}
