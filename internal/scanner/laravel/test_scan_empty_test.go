//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what 빈 디렉토리 스캔 테스트
package laravel

import "testing"

func TestScan_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Endpoints) != 0 {
		t.Errorf("expected 0 endpoints, got %d", len(result.Endpoints))
	}
}
