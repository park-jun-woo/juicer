//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestReadConfigDefault 테스트
package nestjs

import "testing"

func TestReadConfigDefault(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/config/app.config.ts", `export default { apiPrefix: process.env.API_PREFIX || 'api' };`)
	if got := readConfigDefault(dir); got != "api" {
		t.Fatalf("got %q", got)
	}
	if got := readConfigDefault(t.TempDir()); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
