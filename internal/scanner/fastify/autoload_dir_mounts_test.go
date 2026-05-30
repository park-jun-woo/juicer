//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what autoloadDirMounts 테스트
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
	// a .d.ts file -> autoloadFilePrefix returns ok=false -> skipped
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

func TestAutoloadDirMounts_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	mounts := autoloadDirMounts("/app/server.ts", dir, "/api")
	if len(mounts) != 0 {
		t.Fatalf("expected no mounts, got %d", len(mounts))
	}
}

func TestAutoloadDirMounts_PluginRefEqualsPath(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "items.ts")
	os.WriteFile(p, []byte("export default 1;"), 0o644)
	mounts := autoloadDirMounts("/app/server.ts", dir, "")
	if len(mounts) != 1 {
		t.Fatalf("expected 1 mount, got %d", len(mounts))
	}
	if mounts[0].PluginRef != p || mounts[0].FilePath != p {
		t.Fatalf("PluginRef/FilePath should equal path, got %+v", mounts[0])
	}
}
