//ff:func feature=scan type=test control=sequence
//ff:what TestExtractPathString_UnknownCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestExtractPathString_UnknownCov(t *testing.T) {
	_, ok := extractPathString(&ast.Ident{Name: "x"})
	if ok {
		t.Fatal("expected not ok")
	}
}
