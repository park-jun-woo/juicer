//ff:func feature=scan type=test control=sequence
//ff:what TestResolveTargetFilePath 테스트
package gogin

import (
	"go/token"
	"testing"
)

func TestResolveTargetFilePath(t *testing.T) {
	fset := token.NewFileSet()
	f := fset.AddFile("/tmp/project/main.go", -1, 100)
	_ = f

	ctx := &groupArgCtx{
		fset: fset,
		root: "/tmp/project",
	}

	got := resolveTargetFilePath(token.Pos(1), ctx)
	_ = got
}
