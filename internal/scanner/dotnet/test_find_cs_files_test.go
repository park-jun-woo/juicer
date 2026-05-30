//ff:func feature=scan type=test control=iteration dimension=1 topic=dotnet
//ff:what TestFindCsFiles 테스트
package dotnet

import (
	"path/filepath"
	"testing"
)

func TestFindCsFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Controllers/UsersController.cs", "class C {}")
	writeFile(t, dir, "bin/Gen.cs", "class G {}")
	files, err := findCsFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		rel, _ := filepath.Rel(dir, f)
		if rel != filepath.Join("Controllers", "UsersController.cs") {
			t.Errorf("unexpected: %s", rel)
		}
	}
}
