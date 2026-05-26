//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDetectGlobalPrefix_NotFound 테스트
package nestjs

import "testing"

func TestDetectGlobalPrefix_NotFound(t *testing.T) {
	dir := t.TempDir()
	prefix := detectGlobalPrefix(dir)
	if prefix != "" {
		t.Fatalf("expected empty, got %q", prefix)
	}
}
