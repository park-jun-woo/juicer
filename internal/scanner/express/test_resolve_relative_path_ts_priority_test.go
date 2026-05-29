//ff:func feature=scan type=test control=sequence topic=express
//ff:what 동일 base에 .ts/.js 공존 시 resolveRelativePath가 .ts를 우선 선택하는지 검증
package express

import (
	"path/filepath"
	"testing"
)

func TestResolveRelativePath_TSPreferredOverJS(t *testing.T) {
	dir := t.TempDir()
	// both .ts and .js present: .ts wins (first in sourceExtensions)
	writeFile(t, dir, "mod.ts", "export const a = 1;")
	writeFile(t, dir, "mod.js", "module.exports = {};")

	if got := resolveRelativePath(dir, "./mod"); got != filepath.Join(dir, "mod.ts") {
		t.Errorf("expected .ts preferred, got %q", got)
	}
}
