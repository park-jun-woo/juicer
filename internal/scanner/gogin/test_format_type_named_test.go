//ff:func feature=scan type=extract control=sequence
//ff:what TestFormatType_Named 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestFormatType_Named(t *testing.T) {
	pkg := types.NewPackage("example.com/test", "test")
	named := types.NewNamed(
		types.NewTypeName(0, pkg, "MyType", nil),
		types.NewStruct(nil, nil),
		nil,
	)
	got := formatType(named)
	if got != "test.MyType" {
		t.Errorf("expected 'test.MyType', got %q", got)
	}
}
