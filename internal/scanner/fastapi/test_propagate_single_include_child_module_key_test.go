//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestPropagateSingleInclude_ChildModuleKey 테스트
package fastapi

import "testing"

func TestPropagateSingleInclude_ChildModuleKey(t *testing.T) {
	fi := newFileInfoSimple(t, "x=1\n", map[string]string{"app": "/api", "router": "/api"})
	srcFI := newFileInfoSimple(t, "x=1\n", map[string]string{"router": ""})
	srcFI.absPath = "/items.py"
	inc := includeCall{parentVar: "app", childVar: "router", childModule: "items"}
	importMap := map[string]string{"items.router": "/items.py"}
	fileByPath := map[string]*fileInfo{"/items.py": srcFI}
	origSnapshot := map[string]map[string]string{
		fi.absPath:  {"app": "/api", "router": "/api"},
		"/items.py": {"router": ""},
	}
	if !propagateSingleInclude(fi, inc, importMap, fileByPath, origSnapshot) {
		t.Fatalf("expected change via childModule; srcFI.prefixes=%v", srcFI.prefixes)
	}
}
