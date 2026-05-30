//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestPropagateSingleInclude_Unchanged 테스트
package fastapi

import "testing"

func TestPropagateSingleInclude_Unchanged(t *testing.T) {

	fi := newFileInfoSimple(t, "x=1\n", map[string]string{"app": "", "users": ""})
	srcFI := newFileInfoSimple(t, "x=1\n", map[string]string{"users": "/users"})
	srcFI.absPath = "/src.py"
	inc := includeCall{parentVar: "app", childVar: "users"}
	importMap := map[string]string{"users": "/src.py"}
	fileByPath := map[string]*fileInfo{"/src.py": srcFI}
	origSnapshot := map[string]map[string]string{
		fi.absPath: {"app": "", "users": ""},
		"/src.py":  {"users": "/users"},
	}
	if propagateSingleInclude(fi, inc, importMap, fileByPath, origSnapshot) {
		t.Fatalf("expected no change; srcFI.prefixes=%v", srcFI.prefixes)
	}
}
