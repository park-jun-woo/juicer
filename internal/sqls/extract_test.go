//ff:func feature=sql type=parse control=sequence
//ff:what TestExtract 테스트
package sqls

import "testing"

func TestExtract_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	result, err := Extract(dir)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}
