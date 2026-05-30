//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestMergeSingleInclude_LocalExists 테스트
package fastapi

import "testing"

func TestMergeSingleInclude_LocalExists(t *testing.T) {
	fi := newFileInfoSimple(t, "x = 1\n", map[string]string{"users": "/already"})
	inc := includeCall{parentVar: "app", childVar: "users"}
	mergeSingleInclude(fi, inc, map[string]string{"users": "/s"}, map[string]map[string]string{})
	if fi.prefixes["users"] != "/already" {
		t.Fatalf("local prefix should be untouched, got %q", fi.prefixes["users"])
	}
}
