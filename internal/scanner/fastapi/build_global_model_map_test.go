//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what buildGlobalModelMap 테스트
package fastapi

import "testing"

func TestBuildGlobalModelMap(t *testing.T) {
	files := []fileInfo{
		{absPath: "/a.py", models: map[string][]pydanticField{"User": nil}},
		{absPath: "/b.py", models: map[string][]pydanticField{"Order": nil}},
	}
	m := buildGlobalModelMap(files)
	if len(m) != 2 {
		t.Fatalf("expected 2 models, got %d", len(m))
	}
	if m["User"].absPath != "/a.py" {
		t.Fatalf("unexpected User path: %s", m["User"].absPath)
	}
	if m["Order"].absPath != "/b.py" {
		t.Fatalf("unexpected Order path: %s", m["Order"].absPath)
	}

	// empty
	m2 := buildGlobalModelMap(nil)
	if len(m2) != 0 {
		t.Fatalf("expected 0, got %d", len(m2))
	}
}
