//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what buildGlobalPrefixMap: prefix 있는 파일만 매핑 / 빈 prefix 스킵
package fastapi

import "testing"

func TestBuildGlobalPrefixMap(t *testing.T) {
	files := []fileInfo{
		{absPath: "/a.py", prefixes: map[string]string{"router": "/api"}},
		{absPath: "/b.py", prefixes: map[string]string{}}, // empty -> skipped
		{absPath: "/c.py", prefixes: map[string]string{"r1": "/v1", "r2": "/v2"}},
	}
	result := buildGlobalPrefixMap(files)
	if len(result) != 2 {
		t.Fatalf("expected 2 entries, got %d: %v", len(result), result)
	}
	if result["/a.py"]["router"] != "/api" {
		t.Errorf("a.py: %v", result["/a.py"])
	}
	if result["/c.py"]["r2"] != "/v2" {
		t.Errorf("c.py: %v", result["/c.py"])
	}
	if _, ok := result["/b.py"]; ok {
		t.Error("b.py with empty prefixes should be skipped")
	}
}
