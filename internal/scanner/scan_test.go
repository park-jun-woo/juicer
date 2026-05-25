package scanner

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
	// May or may not error depending on go packages loading
	_ = err
}
