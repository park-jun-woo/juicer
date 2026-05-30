//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestParseAllPHPFiles 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestParseAllPHPFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "a.php", "<?php class A {}")
	writeFile(t, dir, "b.php", "<?php class B {}")
	parsed := parseAllPHPFiles(dir, []string{
		filepath.Join(dir, "a.php"),
		filepath.Join(dir, "b.php"),
		filepath.Join(dir, "missing.php"),
	})
	if len(parsed) != 2 {
		t.Fatalf("expected 2, got %d", len(parsed))
	}
}
