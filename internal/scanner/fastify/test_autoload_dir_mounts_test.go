//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestAutoloadDirMounts 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAutoloadDirMounts(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "index.ts"), []byte("export default 1;"), 0o644)
	os.WriteFile(filepath.Join(dir, "users.ts"), []byte("export default 1;"), 0o644)
	sub := filepath.Join(dir, "admin")
	os.MkdirAll(sub, 0o755)
	os.WriteFile(filepath.Join(sub, "index.ts"), []byte("export default 1;"), 0o644)

	os.WriteFile(filepath.Join(dir, "types.d.ts"), []byte("export {};"), 0o644)

	mounts := autoloadDirMounts("/app/server.ts", dir, "/api")
	if len(mounts) != 3 {
		t.Fatalf("expected 3 mounts (.d.ts skipped), got %d: %v", len(mounts), mounts)
	}
	for _, m := range mounts {
		if m.SourceFile != "/app/server.ts" {
			t.Errorf("SourceFile = %q", m.SourceFile)
		}
		if m.Prefix == "" {
			t.Errorf("empty prefix for %s", m.FilePath)
		}
	}
}
