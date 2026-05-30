//ff:func feature=scan type=test control=sequence
//ff:what resolveTargetFilePath — 대상 파일 상대 경로 해석 테스트
package fiber

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestResolveTargetFilePath_Relative(t *testing.T) {
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "/proj/sub/m.go", "package m\n", 0)
	ctx := &groupArgCtx{fset: fset, root: "/proj"}
	if got := resolveTargetFilePath(file.Pos(), ctx); got != "sub/m.go" {
		t.Fatalf("rel path = %q, want sub/m.go", got)
	}
}

func TestResolveTargetFilePath_RelError(t *testing.T) {
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "/proj/m.go", "package m\n", 0)
	// root is relative while path is absolute -> filepath.Rel errors -> absPath
	ctx := &groupArgCtx{fset: fset, root: "relative-root"}
	if got := resolveTargetFilePath(file.Pos(), ctx); got != "/proj/m.go" {
		t.Fatalf("expected absolute fallback, got %q", got)
	}
}
