//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindClassFieldDefaultInFile_ClassMissing 테스트
package fastapi

import "testing"

func TestFindClassFieldDefaultInFile_ClassMissing(t *testing.T) {
	fi := fileInfoFor(t, "class Config:\n    x: str = \"y\"\n")
	if got := findClassFieldDefaultInFile(fi, "Other", "x"); got != "" {
		t.Fatalf("got %q", got)
	}
}
