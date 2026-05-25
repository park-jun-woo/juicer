//ff:func feature=sql type=parse control=sequence
//ff:what TestQueryExists_NotFound 테스트
package sqls

import (
	"os"
	"testing"
)

func TestQueryExists_NotFound(t *testing.T) {
	dir, err := os.MkdirTemp("", "qe-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	if queryExists(dir, "Missing") {
		t.Fatal("expected false")
	}
}
