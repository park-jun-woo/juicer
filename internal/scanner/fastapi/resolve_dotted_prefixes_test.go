//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what resolveDottedPrefixes 테스트
package fastapi

import "testing"

func TestResolveDottedPrefixes(t *testing.T) {
	cfg := mkParsedFile(t, "class Settings:\n    API = \"/v1\"\n\nsettings = Settings()\n")

	router := mkParsedFile(t, "r = 1\n")
	router.prefixes = map[string]string{
		"r":   "settings.API", // dotted -> resolved to "/v1"
		"lit": "/static",       // not dotted -> unchanged
	}

	files := []fileInfo{cfg, router}
	resolveDottedPrefixes("/root", files)

	if files[1].prefixes["r"] != "/v1" {
		t.Errorf("dotted prefix = %q, want %q", files[1].prefixes["r"], "/v1")
	}
	if files[1].prefixes["lit"] != "/static" {
		t.Errorf("literal prefix changed to %q", files[1].prefixes["lit"])
	}
}
