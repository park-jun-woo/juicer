//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFallbackGlobalPrefix 테스트
package nestjs

import "testing"

func TestFallbackGlobalPrefix(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, ".env", "API_PREFIX=fromenv\n")
	if got := fallbackGlobalPrefix(dir); got != "fromenv" {
		t.Fatalf("got %q", got)
	}

	if got := fallbackGlobalPrefix(t.TempDir()); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
