//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestCollectAutoloadMounts_WithRegister 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCollectAutoloadMounts_WithRegister(t *testing.T) {
	root := t.TempDir()

	os.MkdirAll(filepath.Join(root, "routes"), 0o755)
	os.WriteFile(filepath.Join(root, "routes", "users.ts"), []byte("export default 1;"), 0o644)
	os.WriteFile(filepath.Join(root, "routes", "items.ts"), []byte("export default 1;"), 0o644)

	appPath := filepath.Join(root, "app.ts")
	appSrc := `
import Fastify from "fastify";
import autoload from "@fastify/autoload";
import { join } from "path";
const app = Fastify();
app.register(autoload, {
  dir: join(__dirname, "routes"),
  options: { prefix: "/api" }
});
`
	fi := mustParse(t, []byte(appSrc))
	fi.Path = appPath
	parsed := map[string]*fileInfo{appPath: fi}

	mounts := collectAutoloadMounts(parsed, root)
	if len(mounts) != 2 {
		t.Fatalf("expected 2 mounts, got %d: %v", len(mounts), mounts)
	}
}
