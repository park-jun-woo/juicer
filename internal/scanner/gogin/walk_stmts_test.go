//ff:func feature=scan type=test control=sequence
//ff:what TestWalkStmts_Empty 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestWalkStmts_Empty(t *testing.T) {
	var out []scanner.Endpoint
	walkStmts(nil, "gin", "test.go", token.NewFileSet(), map[string]*routerInfo{}, &out, map[int][]ast.Expr{})
	if len(out) != 0 {
		t.Fatal("expected empty")
	}
}

