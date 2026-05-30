//ff:func feature=scan type=test control=sequence topic=actix
//ff:what Scan — 탐색 실패 경계 분기를 검증
package actix

import (
	"path/filepath"
	"testing"
)

func TestScan_BadRoot(t *testing.T) {
	// A non-existent root makes findRsFiles return an error.
	_, err := Scan(filepath.Join(t.TempDir(), "no-such-dir"))
	if err == nil {
		t.Fatal("expected error for non-existent root")
	}
}
