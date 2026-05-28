//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestScan_EmptyDir -- 빈 디렉토리 스캔 테스트
package dotnet

import "testing"

func TestScan_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Endpoints) != 0 {
		t.Errorf("expected 0 endpoints, got %d", len(result.Endpoints))
	}
}
