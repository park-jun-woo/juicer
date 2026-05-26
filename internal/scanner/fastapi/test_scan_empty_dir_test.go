//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestScan_EmptyDir 테스트
package fastapi

import "testing"

func TestScan_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Endpoints) != 0 {
		t.Fatalf("expected 0 endpoints, got %d", len(result.Endpoints))
	}
}
