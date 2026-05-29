//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what laravel이 아닌 composer.json에서 미감지 확인
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectLaravel_NoLaravel(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "composer.json"), []byte(`{
    "require": {
        "php": "^8.1",
        "symfony/framework-bundle": "^6.0"
    }
}`), 0o644)
	if detectLaravel(dir) {
		t.Error("expected detectLaravel = false for non-laravel project")
	}
}
