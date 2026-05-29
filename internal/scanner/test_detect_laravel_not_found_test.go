//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what composer.json이 없으면 laravel 미감지 확인
package scanner

import (
	"testing"
)

func TestDetectLaravel_NotFound(t *testing.T) {
	dir := t.TempDir()
	if detectLaravel(dir) {
		t.Error("expected detectLaravel = false for missing composer.json")
	}
}
