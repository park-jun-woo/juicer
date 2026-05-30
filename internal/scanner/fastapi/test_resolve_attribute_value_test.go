//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestResolveAttributeValue 테스트
package fastapi

import "testing"

func TestResolveAttributeValue(t *testing.T) {
	src := "class Settings:\n    name = \"prod\"\n\nsettings = Settings()\n"
	fi := mkParsedFile(t, src)
	files := []fileInfo{fi}

	if got := resolveAttributeValue("/root", files, "noattr"); got != "" {
		t.Errorf("expected empty for no-dot, got %q", got)
	}

	if got := resolveAttributeValue("/root", []fileInfo{}, "settings.name"); got != "" {
		t.Errorf("expected empty for no files, got %q", got)
	}

	got := resolveAttributeValue("/root", files, "settings.name")
	if got != "prod" {
		t.Errorf("resolveAttributeValue = %q, want %q", got, "prod")
	}
}
