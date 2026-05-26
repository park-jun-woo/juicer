//ff:func feature=scan type=test control=sequence
//ff:what TestResolveTargetFilePath 테스트
package gogin

import (
	"go/token"
	"testing"
)

func TestResolveTargetFilePath_ValidPos(t *testing.T) {
	fset := token.NewFileSet()
	f := fset.AddFile("/tmp/test/main.go", -1, 100)
	_ = f
	ctx := &groupArgCtx{
		fset: fset,
		root: "/tmp/test",
	}
	// pos 1 is in the file we just added
	got := resolveTargetFilePath(1, ctx)
	if got != "main.go" {
		t.Fatalf("expected main.go, got %q", got)
	}
}
