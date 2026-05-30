//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestAutoloadDirMounts_EmptyDir 테스트
package fastify

import "testing"

func TestAutoloadDirMounts_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	mounts := autoloadDirMounts("/app/server.ts", dir, "/api")
	if len(mounts) != 0 {
		t.Fatalf("expected no mounts, got %d", len(mounts))
	}
}
