//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestScan_Empty 테스트
package fastify

import "testing"

func TestScan_Empty(t *testing.T) {
	dir := t.TempDir()
	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 0 {
		t.Errorf("expected 0 endpoints, got %d", len(result.Endpoints))
	}
}
