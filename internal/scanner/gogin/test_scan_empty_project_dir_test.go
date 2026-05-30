//ff:func feature=scan type=test control=sequence
//ff:what TestScan_EmptyProjectDir 테스트
package gogin

import (
	"os"
	"testing"
)

func TestScan_EmptyProjectDir(t *testing.T) {
	dir, err := os.MkdirTemp("", "scan-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	_, err = Scan(dir)

	_ = err
}
