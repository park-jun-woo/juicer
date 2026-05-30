//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what resolveChildModuleImport 테스트
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveChildModuleImport(t *testing.T) {
	root := t.TempDir()

	// absolute layout: <root>/app/routes/users.py
	os.MkdirAll(filepath.Join(root, "app", "routes"), 0o755)
	abs := filepath.Join(root, "app", "routes", "users.py")
	os.WriteFile(abs, []byte(""), 0o644)

	// relative layout next to referrer: <referrer>/pkg/items.py
	referrer := filepath.Join(root, "ref")
	os.MkdirAll(filepath.Join(referrer, "pkg"), 0o755)
	rel := filepath.Join(referrer, "pkg", "items.py")
	os.WriteFile(rel, []byte(""), 0o644)

	// no matching import
	if got := resolveChildModuleImport(root, referrer, "users", nil); got != "" {
		t.Errorf("expected empty for no imports, got %q", got)
	}

	// absolute fallback: relative resolution fails, absolute succeeds
	impsAbs := []importInfo{{name: "users", module: "app.routes"}}
	if got := resolveChildModuleImport(root, referrer, "users", impsAbs); got != abs {
		t.Errorf("abs resolve = %q, want %q", got, abs)
	}

	// relative resolution: module ".pkg" + ".items"
	impsRel := []importInfo{{name: "items", module: ".pkg"}}
	if got := resolveChildModuleImport(root, referrer, "items", impsRel); got != rel {
		t.Errorf("rel resolve = %q, want %q", got, rel)
	}
}
