//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestResolveConstantValue 테스트
package spring

import (
	"path/filepath"
	"testing"
)

func TestResolveConstantValue(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Const.java", `class Const { public static final String LIMIT = "10"; }`)
	got := resolveConstantValue("Const.LIMIT", map[string]string{}, filepath.Join(dir, "R.java"), dir)
	if got != "10" {
		t.Fatalf("got %q", got)
	}

	if got := resolveConstantValue("Unknown.X", map[string]string{}, filepath.Join(dir, "R.java"), dir); got != "Unknown.X" {
		t.Fatalf("unresolved: %q", got)
	}

	if got := resolveConstantValue("plain", map[string]string{}, "", ""); got != "plain" {
		t.Fatalf("plain: %q", got)
	}
}
