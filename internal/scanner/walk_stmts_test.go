//ff:func feature=scan type=extract control=sequence
//ff:what TestWalkStmts_Empty 테스트
package scanner

import (
	"go/token"
	"testing"
)

func TestWalkStmts_Empty(t *testing.T) {
	var out []Endpoint
	walkStmts(nil, "gin", "test.go", token.NewFileSet(), map[string]*routerInfo{}, &out)
	if len(out) != 0 {
		t.Fatal("expected empty")
	}
}
