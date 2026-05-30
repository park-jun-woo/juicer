//ff:func feature=scan type=test control=sequence
//ff:what TestResolveTargetFilePath_Relative 테스트
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
