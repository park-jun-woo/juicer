//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what mergeSingleInclude: 병합 / 로컬존재 스킵 / srcFile없음 / srcPrefixes없음 / childModule키
package fastapi

import "testing"

func newFileInfoSimple(t *testing.T, src string, prefixes map[string]string) *fileInfo {
	t.Helper()
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	return &fileInfo{absPath: "/main.py", src: []byte(src), root: root, prefixes: prefixes}
}

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

func TestMergeSingleInclude_LocalExists(t *testing.T) {
	fi := newFileInfoSimple(t, "x = 1\n", map[string]string{"users": "/already"})
	inc := includeCall{parentVar: "app", childVar: "users"}
	mergeSingleInclude(fi, inc, map[string]string{"users": "/s"}, map[string]map[string]string{})
	if fi.prefixes["users"] != "/already" {
		t.Fatalf("local prefix should be untouched, got %q", fi.prefixes["users"])
	}
}

func TestMergeSingleInclude_SrcFileMissing(t *testing.T) {
	fi := newFileInfoSimple(t, "x = 1\n", map[string]string{"app": ""})
	inc := includeCall{parentVar: "app", childVar: "users"}
	mergeSingleInclude(fi, inc, map[string]string{}, map[string]map[string]string{})
	if _, ok := fi.prefixes["users"]; ok {
		t.Fatalf("should not merge when srcFile missing")
	}
}

func TestMergeSingleInclude_NoSrcPrefixes(t *testing.T) {
	fi := newFileInfoSimple(t, "x = 1\n", map[string]string{"app": ""})
	inc := includeCall{parentVar: "app", childVar: "users"}
	importMap := map[string]string{"users": "/src/users.py"}
	mergeSingleInclude(fi, inc, importMap, map[string]map[string]string{})
	if _, ok := fi.prefixes["users"]; ok {
		t.Fatalf("should not merge when src prefixes nil")
	}
}

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
