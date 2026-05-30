//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what resolveDottedPrefixesForFile 테스트
package fastapi

import "testing"

func TestResolveDottedPrefixesForFile(t *testing.T) {
	cfg := mkParsedFile(t, "class Settings:\n    API = \"/v1\"\n\nsettings = Settings()\n")

	router := mkParsedFile(t, "r = 1\n")
	router.prefixes = map[string]string{
		"resolved":   "/settings.API",  // leading slash trimmed, dotted -> "/v1"
		"unresolved": "settings.NOPE",  // dotted but attr missing -> unchanged
		"plain":      "/static",        // not dotted -> continue
	}

	files := []fileInfo{cfg, router}
	resolveDottedPrefixesForFile("/root", files, 1)

	if files[1].prefixes["resolved"] != "/v1" {
		t.Errorf("resolved = %q, want %q", files[1].prefixes["resolved"], "/v1")
	}
	if files[1].prefixes["unresolved"] != "settings.NOPE" {
		t.Errorf("unresolved changed to %q", files[1].prefixes["unresolved"])
	}
	if files[1].prefixes["plain"] != "/static" {
		t.Errorf("plain changed to %q", files[1].prefixes["plain"])
	}
}
