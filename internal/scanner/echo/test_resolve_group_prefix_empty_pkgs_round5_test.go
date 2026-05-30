//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestResolveGroupPrefix_EmptyPkgs_Round5 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestResolveGroupPrefix_EmptyPkgs_Round5(t *testing.T) {

	resolveGroupPrefix(nil, "/root", buildFuncIndex(nil), nil, map[int][]ast.Expr{})
}
