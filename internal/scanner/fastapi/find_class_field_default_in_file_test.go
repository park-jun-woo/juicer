//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what findClassFieldDefaultInFile: 클래스/필드 기본값 / 클래스 미존재
package fastapi

import "testing"

func fileInfoFor(t *testing.T, src string) fileInfo {
	t.Helper()
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	return fileInfo{src: []byte(src), root: root}
}

func TestFindClassFieldDefaultInFile_Found(t *testing.T) {
	fi := fileInfoFor(t, "class Config:\n    table_name: str = \"users\"\n")
	if got := findClassFieldDefaultInFile(fi, "Config", "table_name"); got != "users" {
		t.Fatalf("got %q", got)
	}
}

func TestFindClassFieldDefaultInFile_ClassMissing(t *testing.T) {
	fi := fileInfoFor(t, "class Config:\n    x: str = \"y\"\n")
	if got := findClassFieldDefaultInFile(fi, "Other", "x"); got != "" {
		t.Fatalf("got %q", got)
	}
}
