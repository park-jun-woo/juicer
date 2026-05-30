//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestMergeSingleInclude_SrcFileMissing 테스트
package fastapi

import "testing"

func TestMergeSingleInclude_SrcFileMissing(t *testing.T) {
	fi := newFileInfoSimple(t, "x = 1\n", map[string]string{"app": ""})
	inc := includeCall{parentVar: "app", childVar: "users"}
	mergeSingleInclude(fi, inc, map[string]string{}, map[string]map[string]string{})
	if _, ok := fi.prefixes["users"]; ok {
		t.Fatalf("should not merge when srcFile missing")
	}
}
