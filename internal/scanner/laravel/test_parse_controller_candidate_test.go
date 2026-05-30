//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestParseControllerCandidate 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestParseControllerCandidate(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "C.php", "<?php class C {}")
	if parseControllerCandidate(dir, filepath.Join(dir, "C.php")) == nil {
		t.Fatal("expected fileInfo")
	}
	if parseControllerCandidate(dir, filepath.Join(dir, "missing.php")) != nil {
		t.Fatal("expected nil for missing")
	}
}
