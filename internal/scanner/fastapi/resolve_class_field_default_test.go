//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what findClassFieldDefault 테스트
package fastapi

import "testing"

func TestFindClassFieldDefault(t *testing.T) {
	fi := mkParsedFile(t, "class Settings:\n    name: str = \"prod\"\n")
	files := []fileInfo{fi}

	// found
	if got := findClassFieldDefault(files, "Settings", "name"); got != "prod" {
		t.Errorf("findClassFieldDefault = %q, want %q", got, "prod")
	}

	// not found -> ""
	if got := findClassFieldDefault(files, "Settings", "missing"); got != "" {
		t.Errorf("expected empty for missing attr, got %q", got)
	}
	if got := findClassFieldDefault(nil, "Settings", "name"); got != "" {
		t.Errorf("expected empty for no files, got %q", got)
	}
}
