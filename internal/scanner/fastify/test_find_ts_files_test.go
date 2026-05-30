//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestFindTSFiles 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindTSFiles(t *testing.T) {
	root := t.TempDir()

	os.WriteFile(filepath.Join(root, "app.ts"), []byte(""), 0o644)
	os.MkdirAll(filepath.Join(root, "src"), 0o755)
	os.WriteFile(filepath.Join(root, "src", "routes.ts"), []byte(""), 0o644)

	os.WriteFile(filepath.Join(root, "types.d.ts"), []byte(""), 0o644)
	os.WriteFile(filepath.Join(root, "x.js"), []byte(""), 0o644)
	os.MkdirAll(filepath.Join(root, "node_modules", "lib"), 0o755)
	os.WriteFile(filepath.Join(root, "node_modules", "lib", "index.ts"), []byte(""), 0o644)
	os.WriteFile(filepath.Join(root, "app.test.ts"), []byte(""), 0o644)

	for _, d := range []string{"dist", "build", ".git"} {
		os.MkdirAll(filepath.Join(root, d), 0o755)
		os.WriteFile(filepath.Join(root, d, "gen.ts"), []byte(""), 0o644)
	}

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
