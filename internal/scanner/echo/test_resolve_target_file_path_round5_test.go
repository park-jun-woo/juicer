//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestResolveTargetFilePath_Round5 테스트
package echo

import (
	"path/filepath"
	"testing"
)

func TestResolveTargetFilePath_Round5(t *testing.T) {
	ctx := emptyGroupCtx()

	f := ctx.fset.AddFile("/root/sub/m.go", -1, 10)
	pos := f.Pos(0)
	got := resolveTargetFilePath(pos, ctx)
	if got != filepath.Join("sub", "m.go") {
		t.Fatalf("rel path: %q", got)
	}
}
