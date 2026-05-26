//ff:func feature=scan type=test control=sequence
//ff:what TestFormatType_NamedCov 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestFormatType_NamedCov(t *testing.T) {
	pkg := types.NewPackage("example.com/pkg", "pkg")
	tn := types.NewTypeName(0, pkg, "MyType", nil)
	named := types.NewNamed(tn, types.Typ[types.Int], nil)
	got := formatType(named)
	if got != "pkg.MyType" {
		t.Fatalf("expected pkg.MyType, got %s", got)
	}
}
