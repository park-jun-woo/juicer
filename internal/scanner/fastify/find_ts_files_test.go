//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what findTSFiles 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindTSFiles(t *testing.T) {
	root := t.TempDir()
	// included
	os.WriteFile(filepath.Join(root, "app.ts"), []byte(""), 0o644)
	os.MkdirAll(filepath.Join(root, "src"), 0o755)
	os.WriteFile(filepath.Join(root, "src", "routes.ts"), []byte(""), 0o644)
	// excluded: .d.ts, .js, node_modules, test file
	os.WriteFile(filepath.Join(root, "types.d.ts"), []byte(""), 0o644)
	os.WriteFile(filepath.Join(root, "x.js"), []byte(""), 0o644)
	os.MkdirAll(filepath.Join(root, "node_modules", "lib"), 0o755)
	os.WriteFile(filepath.Join(root, "node_modules", "lib", "index.ts"), []byte(""), 0o644)
	os.WriteFile(filepath.Join(root, "app.test.ts"), []byte(""), 0o644)
	// excluded dirs: dist, build, .git
	for _, d := range []string{"dist", "build", ".git"} {
		os.MkdirAll(filepath.Join(root, d), 0o755)
		os.WriteFile(filepath.Join(root, d, "gen.ts"), []byte(""), 0o644)
	}
	// excluded test directory
	os.MkdirAll(filepath.Join(root, "__tests__"), 0o755)
	os.WriteFile(filepath.Join(root, "__tests__", "spec.ts"), []byte(""), 0o644)

	files, err := findTSFiles(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 2 {
		t.Fatalf("expected 2 .ts files, got %d: %v", len(files), files)
	}
}

func TestFindTSFiles_Error(t *testing.T) {
	_, err := findTSFiles("/no/such/path/zzz")
	if err == nil {
		t.Fatal("expected error for missing root")
	}
}
