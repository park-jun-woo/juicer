//ff:func feature=scan type=test control=sequence
//ff:what resolveGroupPrefix 전 분기 테스트
package gogin

import (
	"go/ast"
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
	"golang.org/x/tools/go/packages"
)

func TestResolveGroupPrefix(t *testing.T) {
	// empty pkgs
	resolveGroupPrefix(nil, "/tmp", &funcIndex{}, nil, nil)

	// pkg with nil TypesInfo
	pkg := &packages.Package{}
	resolveGroupPrefix([]*packages.Package{pkg}, "/tmp", &funcIndex{}, nil, nil)

	// pkg with TypesInfo but no syntax
	pkg2 := &packages.Package{
		TypesInfo: nil,
	}
	resolveGroupPrefix([]*packages.Package{pkg2}, "/tmp", &funcIndex{}, []scanner.Endpoint{}, map[int][]ast.Expr{})
}
