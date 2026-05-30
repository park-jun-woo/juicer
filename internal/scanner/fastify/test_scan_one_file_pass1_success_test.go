//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestScanOneFilePass1_Success 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanOneFilePass1_Success(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "app.ts")
	os.WriteFile(p, []byte(`
import Fastify from "fastify";
const app = Fastify();
app.get("/x", h);
`), 0o644)
	res := scanOneFilePass1(p, dir)
	if res == nil {
		t.Fatal("expected non-nil result")
	}
	if res.fi == nil || !res.instances["app"] {
		t.Fatalf("expected app instance, got %+v", res.instances)
	}
}
