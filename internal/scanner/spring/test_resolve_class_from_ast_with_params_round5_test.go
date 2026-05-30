//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestResolveClassFromASTWithParams_Round5 테스트
package spring

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveClassFromASTWithParams_Round5(t *testing.T) {
	root, src := sParse(t, `class Page<T> { public T content; public int total; }`)
	fields, params := resolveClassFromASTWithParams(root, src, "Page", "C.java", "/abs", map[string][]scanner.Field{})
	if len(fields) == 0 {
		t.Fatalf("fields empty")
	}
	_ = params
}
