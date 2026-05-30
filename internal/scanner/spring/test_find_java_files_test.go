//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestFindJavaFiles 테스트
package spring

import "testing"

func TestFindJavaFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/Foo.java", "class Foo {}")
	writeFile(t, dir, "target/Bar.java", "class Bar {}")
	files, err := findJavaFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 {
		t.Fatalf("expected 1, got %v", files)
	}
}
