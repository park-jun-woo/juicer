//ff:func feature=scan type=test control=sequence
//ff:what TestResolveTargetGinAlias 테스트
package gogin

import (
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestResolveTargetGinAlias(t *testing.T) {
	fset := token.NewFileSet()
	ctx := &groupArgCtx{
		fset: fset,
		pkgs: []*packages.Package{},
	}

	got := resolveTargetGinAlias(token.Pos(1), ctx)
	if got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
