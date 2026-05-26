//ff:func feature=scan type=test control=sequence
//ff:what TestBuildFuncIndex_NilInfoBranch 테스트
package gogin

import (
	"testing"
	"golang.org/x/tools/go/packages"
)

func TestBuildFuncIndex_NilInfoBranch(t *testing.T) {
	pkg := &packages.Package{TypesInfo: nil}
	idx := buildFuncIndex([]*packages.Package{pkg})
	if len(idx.byPos) != 0 {
		t.Fatal("expected empty index")
	}
}
