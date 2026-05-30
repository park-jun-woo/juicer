//ff:func feature=scan type=test control=iteration dimension=1 topic=quarkus
//ff:what TestFindJavaFiles 테스트
package quarkus

import (
	"path/filepath"
	"testing"
)

func TestFindJavaFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/Foo.java", "class Foo {}")
	writeFile(t, dir, "target/Bar.java", "class Bar {}")
	files, err := findJavaFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		rel, _ := filepath.Rel(dir, f)
		if rel != filepath.Join("src", "Foo.java") {
			t.Errorf("unexpected: %s", rel)
		}
	}
}
