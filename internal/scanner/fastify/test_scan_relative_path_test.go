//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestScan_RelativePath 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_RelativePath(t *testing.T) {

	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "app.ts"), []byte(`
import Fastify from "fastify";
const app = Fastify();
app.get("/health", health);
`), 0o644)
	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Endpoints) != 1 || result.Endpoints[0].Path != "/health" {
		t.Fatalf("expected 1 /health endpoint, got %+v", result.Endpoints)
	}
}
