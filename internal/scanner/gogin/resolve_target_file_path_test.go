//ff:func feature=scan type=test control=sequence
//ff:what resolveTargetFilePath 전 분기 테스트
package gogin

import (
	"go/token"
	rtfpPars "go/parser"
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

	// valid pos
	got := resolveTargetFilePath(token.Pos(1), ctx)
	_ = got
}

func TestResolveTargetFilePath_RelError(t *testing.T) {
	fset := token.NewFileSet()
	file, _ := rtfpPars.ParseFile(fset, "/proj/m.go", "package m\n", 0)
	// relative root + absolute path -> filepath.Rel errors -> returns absPath
	ctx := &groupArgCtx{fset: fset, root: "relative-root"}
	if got := resolveTargetFilePath(file.Pos(), ctx); got != "/proj/m.go" {
		t.Fatalf("expected absolute fallback, got %q", got)
	}
}
