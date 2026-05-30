//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what resolveAbsoluteImportPath 테스트
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveAbsoluteImportPath(t *testing.T) {
	root := t.TempDir()

	// app/sneakers.py
	os.MkdirAll(filepath.Join(root, "app"), 0o755)
	pyFile := filepath.Join(root, "app", "sneakers.py")
	os.WriteFile(pyFile, []byte("class S: pass"), 0o644)

	// pkg/__init__.py
	os.MkdirAll(filepath.Join(root, "pkg"), 0o755)
	initFile := filepath.Join(root, "pkg", "__init__.py")
	os.WriteFile(initFile, []byte(""), 0o644)

	tests := []struct {
		name   string
		module string
		want   string
	}{
		{"empty module", "", ""},
		{"relative module", ".models", ""},
		{"py file", "app.sneakers", pyFile},
		{"init file", "pkg", initFile},
		{"unresolvable", "does.not.exist", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resolveAbsoluteImportPath(root, tt.module); got != tt.want {
				t.Errorf("resolveAbsoluteImportPath(%q) = %q, want %q", tt.module, got, tt.want)
			}
		})
	}
}
