//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what TestFindPHPFiles 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestFindPHPFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app/Foo.php", "<?php")
	writeFile(t, dir, "vendor/lib.php", "<?php")
	writeFile(t, dir, "Foo.test.php", "<?php")
	files, err := findPHPFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		rel, _ := filepath.Rel(dir, f)
		if rel != filepath.Join("app", "Foo.php") {
			t.Errorf("unexpected file: %s", rel)
		}
	}
}
