//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what resolvePluginFilePaths 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolvePluginFilePaths(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "items.ts"), []byte(""), 0o644)

	mounts := []pluginMount{
		{PluginRef: "usersPlugin"},          // resolved via imports map
		{PluginRef: "./items"},              // resolved via relative path
		{PluginRef: inlineRef},              // skipped (inline)
		{PluginRef: ""},                     // skipped (empty)
		{PluginRef: "unresolvable-extern"},  // not resolved -> FilePath stays ""
	}
	imports := map[string]string{"usersPlugin": "/abs/users.ts"}

	resolvePluginFilePaths(mounts, imports, dir)

	if mounts[0].FilePath != "/abs/users.ts" {
		t.Errorf("import-resolved FilePath = %q", mounts[0].FilePath)
	}
	if mounts[1].FilePath != filepath.Join(dir, "items.ts") {
		t.Errorf("relative-resolved FilePath = %q", mounts[1].FilePath)
	}
	if mounts[2].FilePath != "" || mounts[3].FilePath != "" {
		t.Errorf("inline/empty should not resolve: %v %v", mounts[2], mounts[3])
	}
	if mounts[4].FilePath != "" {
		t.Errorf("unresolvable should stay empty, got %q", mounts[4].FilePath)
	}
}
