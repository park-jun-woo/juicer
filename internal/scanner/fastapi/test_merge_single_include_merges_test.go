//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestMergeSingleInclude_Merges 테스트
package fastapi

import "testing"

func TestMergeSingleInclude_Merges(t *testing.T) {
	fi := newFileInfoSimple(t, "x = 1\n", map[string]string{"app": "/v1"})
	inc := includeCall{parentVar: "app", childVar: "users"}
	importMap := map[string]string{"users": "/src/users.py"}
	global := map[string]map[string]string{"/src/users.py": {"users": "/u"}}
	mergeSingleInclude(fi, inc, importMap, global)
	if got := fi.prefixes["users"]; got != "/v1/u" {
		t.Fatalf("got %q", got)
	}
	if global["/main.py"]["users"] != "/v1/u" {
		t.Fatalf("global not updated: %v", global["/main.py"])
	}
}
