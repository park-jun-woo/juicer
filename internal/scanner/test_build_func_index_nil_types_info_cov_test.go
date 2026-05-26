//ff:func feature=scan type=test control=sequence
//ff:what TestBuildFuncIndex_NilTypesInfoCov 테스트
package scanner

import (
	"testing"
	"golang.org/x/tools/go/packages"
)

func TestBuildFuncIndex_NilTypesInfoCov(t *testing.T) {
	pkg := &packages.Package{}
	idx := buildFuncIndex([]*packages.Package{pkg})
	if idx == nil {
		t.Fatal("expected non-nil")
	}
	if len(idx.byPos) != 0 {
		t.Fatal("expected empty index")
	}
}
