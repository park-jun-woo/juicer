//ff:func feature=scan type=test control=sequence
//ff:what TestResolveTargetFilePath_RelError 테스트
package fiber

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestResolveTargetFilePath_RelError(t *testing.T) {
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "/proj/m.go", "package m\n", 0)

	ctx := &groupArgCtx{fset: fset, root: "relative-root"}
	if got := resolveTargetFilePath(file.Pos(), ctx); got != "/proj/m.go" {
		t.Fatalf("expected absolute fallback, got %q", got)
	}
}
