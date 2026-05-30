//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what propagateSingleInclude: srcFile없음 / srcFI없음 / origVar없음 / 변경
package fastapi

import "testing"

func TestPropagateSingleInclude_SrcFileMissing(t *testing.T) {
	fi := newFileInfoSimple(t, "x=1\n", map[string]string{})
	inc := includeCall{parentVar: "app", childVar: "users"}
	if propagateSingleInclude(fi, inc, map[string]string{}, map[string]*fileInfo{}, map[string]map[string]string{}) {
		t.Fatal("expected false when srcFile missing")
	}
}

func TestPropagateSingleInclude_SrcFINil(t *testing.T) {
	fi := newFileInfoSimple(t, "x=1\n", map[string]string{})
	inc := includeCall{parentVar: "app", childVar: "users"}
	importMap := map[string]string{"users": "/src.py"}
	// fileByPath has no entry for /src.py -> srcFI nil
	if propagateSingleInclude(fi, inc, importMap, map[string]*fileInfo{}, map[string]map[string]string{}) {
		t.Fatal("expected false when srcFI nil")
	}
}

func TestPropagateSingleInclude_OrigVarMissing(t *testing.T) {
	fi := newFileInfoSimple(t, "x=1\n", map[string]string{})
	srcFI := newFileInfoSimple(t, "x=1\n", map[string]string{}) // no "users" prefix
	srcFI.absPath = "/src.py"
	inc := includeCall{parentVar: "app", childVar: "users"}
	importMap := map[string]string{"users": "/src.py"}
	fileByPath := map[string]*fileInfo{"/src.py": srcFI}
	if propagateSingleInclude(fi, inc, importMap, fileByPath, map[string]map[string]string{}) {
		t.Fatal("expected false when origVar missing")
	}
}

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

func TestPropagateSingleInclude_Unchanged(t *testing.T) {
	// accumulated equals srcFI's current prefix -> no change, returns false
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
