//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestMergeSingleInclude_NoSrcPrefixes 테스트
package fastapi

import "testing"

func TestMergeSingleInclude_NoSrcPrefixes(t *testing.T) {
	fi := newFileInfoSimple(t, "x = 1\n", map[string]string{"app": ""})
	inc := includeCall{parentVar: "app", childVar: "users"}
	importMap := map[string]string{"users": "/src/users.py"}
	mergeSingleInclude(fi, inc, importMap, map[string]map[string]string{})
	if _, ok := fi.prefixes["users"]; ok {
		t.Fatalf("should not merge when src prefixes nil")
	}
}
