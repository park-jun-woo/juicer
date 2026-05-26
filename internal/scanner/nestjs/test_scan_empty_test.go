//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what NestJS 빈 src 디렉토리 스캔 테스트
package nestjs

import "testing"

func TestScan_EmptySrcDir(t *testing.T) {
	dir := t.TempDir()
	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Endpoints) != 0 {
		t.Errorf("expected 0 endpoints, got %d", len(result.Endpoints))
	}
}
