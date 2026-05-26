//ff:func feature=scan type=test control=sequence
//ff:what resolveTargetGinAlias 전 분기 테스트
package gogin

import (
	"go/token"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestResolveTargetGinAlias(t *testing.T) {
	fset := token.NewFileSet()
	ctx := &groupArgCtx{
		fset: fset,
		pkgs: []*packages.Package{},
	}

	// no matching file -> empty
	got := resolveTargetGinAlias(token.Pos(1), ctx)
	if got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
