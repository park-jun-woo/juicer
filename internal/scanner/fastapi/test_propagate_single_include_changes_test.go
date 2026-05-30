//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestPropagateSingleInclude_Changes 테스트
package fastapi

import "testing"

func TestPropagateSingleInclude_Changes(t *testing.T) {
	fi := newFileInfoSimple(t, "x=1\n", map[string]string{"app": "/api", "users": "/api"})
	srcFI := newFileInfoSimple(t, "x=1\n", map[string]string{"users": ""})
	srcFI.absPath = "/src.py"
	inc := includeCall{parentVar: "app", childVar: "users"}
	importMap := map[string]string{"users": "/src.py"}
	fileByPath := map[string]*fileInfo{"/src.py": srcFI}
	origSnapshot := map[string]map[string]string{
		fi.absPath: {"app": "/api", "users": "/api"},
		"/src.py":  {"users": ""},
	}
	if !propagateSingleInclude(fi, inc, importMap, fileByPath, origSnapshot) {
		t.Fatalf("expected change; srcFI.prefixes=%v", srcFI.prefixes)
	}
}
