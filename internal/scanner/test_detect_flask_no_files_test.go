//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFlask_NoFiles 테스트
package scanner

import "testing"

func TestDetectFlask_NoFiles(t *testing.T) {
	if detectFlask(t.TempDir()) {
		t.Fatal("expected false when no dependency files")
	}
}
