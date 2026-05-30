//ff:func feature=scan type=test control=sequence
//ff:what TestDetectQuarkus_NoFiles 테스트
package scanner

import "testing"

func TestDetectQuarkus_NoFiles(t *testing.T) {
	if detectQuarkus(t.TempDir()) {
		t.Fatal("expected false when no build files")
	}
}
