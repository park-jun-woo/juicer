//ff:func feature=scan type=test control=sequence topic=express
//ff:what resolveRelativePath가 .js 파일과 index.js 디렉터리 진입점을 해석하는지 검증
package express

import (
	"path/filepath"
	"testing"
)

func TestResolveRelativePath_JS(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "auth.route.js", "module.exports = {};")
	writeFile(t, dir, "routes/v1/index.js", "module.exports = {};")

	// ./auth.route → auth.route.js
	if got := resolveRelativePath(dir, "./auth.route"); got != filepath.Join(dir, "auth.route.js") {
		t.Errorf("auth.route: got %q", got)
	}
	// ./routes/v1 → routes/v1/index.js
	if got := resolveRelativePath(dir, "./routes/v1"); got != filepath.Join(dir, "routes/v1/index.js") {
		t.Errorf("routes/v1: got %q", got)
	}
	// non-relative import → empty
	if got := resolveRelativePath(dir, "express"); got != "" {
		t.Errorf("bare import should not resolve, got %q", got)
	}
}
