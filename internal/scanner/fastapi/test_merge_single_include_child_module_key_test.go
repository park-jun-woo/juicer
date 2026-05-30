//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestMergeSingleInclude_ChildModuleKey 테스트
package fastapi

import "testing"

func TestMergeSingleInclude_ChildModuleKey(t *testing.T) {
	fi := newFileInfoSimple(t, "x = 1\n", map[string]string{"app": ""})
	inc := includeCall{parentVar: "app", childVar: "router", childModule: "items"}
	importMap := map[string]string{"items.router": "/src/items.py"}
	global := map[string]map[string]string{"/src/items.py": {"router": "/items"}}
	mergeSingleInclude(fi, inc, importMap, global)
	if _, ok := fi.prefixes["router"]; !ok {
		t.Fatalf("expected merge via childModule key, got %v", fi.prefixes)
	}
}
