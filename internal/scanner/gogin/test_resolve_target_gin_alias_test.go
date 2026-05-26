//ff:func feature=scan type=test control=sequence
//ff:what TestResolveTargetGinAlias 테스트
package gogin

import (
	"go/token"
	"testing"
)

func TestResolveTargetGinAlias_NoFile(t *testing.T) {
	ctx := &groupArgCtx{
		pkgs: nil,
	}
	if got := resolveTargetGinAlias(token.NoPos, ctx); got != "" {
		t.Fatalf("expected empty for no matching file, got %q", got)
	}
}
