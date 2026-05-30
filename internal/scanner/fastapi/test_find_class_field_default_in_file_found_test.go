//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindClassFieldDefaultInFile_Found 테스트
package fastapi

import "testing"

func TestFindClassFieldDefaultInFile_Found(t *testing.T) {
	fi := fileInfoFor(t, "class Config:\n    table_name: str = \"users\"\n")
	if got := findClassFieldDefaultInFile(fi, "Config", "table_name"); got != "users" {
		t.Fatalf("got %q", got)
	}
}
