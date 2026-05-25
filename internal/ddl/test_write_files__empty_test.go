//ff:func feature=ddl type=render control=sequence
//ff:what TestWriteFiles_Empty 테스트
package ddl

import "testing"

func TestWriteFiles_Empty(t *testing.T) {
	dir := t.TempDir()
	if err := WriteFiles(map[string]*Table{}, dir); err != nil {
		t.Fatal(err)
	}
}
