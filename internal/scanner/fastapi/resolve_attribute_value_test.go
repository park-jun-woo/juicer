//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what resolveAttributeValue 테스트
package fastapi

import "testing"

func mkParsedFile(t *testing.T, src string) fileInfo {
	t.Helper()
	b := []byte(src)
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	return fileInfo{src: b, root: root}
}

func TestResolveAttributeValue(t *testing.T) {
	src := "class Settings:\n    name = \"prod\"\n\nsettings = Settings()\n"
	fi := mkParsedFile(t, src)
	files := []fileInfo{fi}

	// no dot -> ""
	if got := resolveAttributeValue("/root", files, "noattr"); got != "" {
		t.Errorf("expected empty for no-dot, got %q", got)
	}

	// objName not assigned to a class anywhere -> "" (continue then fallthrough)
	if got := resolveAttributeValue("/root", []fileInfo{}, "settings.name"); got != "" {
		t.Errorf("expected empty for no files, got %q", got)
	}

	// resolves default value
	got := resolveAttributeValue("/root", files, "settings.name")
	if got != "prod" {
		t.Errorf("resolveAttributeValue = %q, want %q", got, "prod")
	}
}
