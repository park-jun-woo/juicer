//ff:func feature=scan type=test control=sequence
//ff:what TestResolveGroupPrefix 테스트
package gogin

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestResolveGroupPrefix(t *testing.T) {

	resolveGroupPrefix(nil, "/tmp", &funcIndex{}, nil, nil)

	pkg := &packages.Package{}
	resolveGroupPrefix([]*packages.Package{pkg}, "/tmp", &funcIndex{}, nil, nil)

	pkg2 := &packages.Package{
		TypesInfo: nil,
	}
	resolveGroupPrefix([]*packages.Package{pkg2}, "/tmp", &funcIndex{}, []scanner.Endpoint{}, map[int][]ast.Expr{})
}
