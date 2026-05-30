//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestAutoloadDirMounts_PluginRefEqualsPath 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

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
