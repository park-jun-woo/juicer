//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestDetectDjango_NoFiles 테스트
package scanner

import "testing"

func TestDetectDjango_NoFiles(t *testing.T) {
	if detectDjango(t.TempDir()) {
		t.Fatal("expected false when no dependency files")
	}
}
