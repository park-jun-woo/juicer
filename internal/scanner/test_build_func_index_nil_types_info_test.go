//ff:func feature=scan type=extract control=sequence
//ff:what TestBuildFuncIndex_NilTypesInfo 테스트
package scanner

import (
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestBuildFuncIndex_NilTypesInfo(t *testing.T) {
	pkg := &packages.Package{
		TypesInfo: nil,
	}
	idx := buildFuncIndex([]*packages.Package{pkg})
	if len(idx.byPos) != 0 {
		t.Errorf("expected 0 functions, got %d", len(idx.byPos))
	}
}
