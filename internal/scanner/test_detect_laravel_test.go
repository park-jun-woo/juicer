//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what composer.json에서 laravel/framework 감지 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectLaravel(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "composer.json"), []byte(`{
    "require": {
        "php": "^8.1",
        "laravel/framework": "^10.0"
    }
}`), 0o644)
	if !detectLaravel(dir) {
		t.Error("expected detectLaravel = true")
	}
}

func TestDetectLaravel_NotFound(t *testing.T) {
	dir := t.TempDir()
	if detectLaravel(dir) {
		t.Error("expected detectLaravel = false for missing composer.json")
	}
}

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
