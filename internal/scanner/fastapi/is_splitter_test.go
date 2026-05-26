//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what isSplitter 테스트
package fastapi

import "testing"

func TestIsSplitter(t *testing.T) {
	if !isSplitter(',') {
		t.Fatal("comma should be splitter")
	}
	if !isSplitter('|') {
		t.Fatal("pipe should be splitter")
	}
	if isSplitter('a') {
		t.Fatal("'a' should not be splitter")
	}
}
