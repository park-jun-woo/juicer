//ff:func feature=scan type=test control=sequence topic=express
//ff:what resolvePathAlias: alias매칭+파일존재 / prefix불일치 / 파일없음
package express

import "testing"

func TestResolvePathAlias_Match(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/api/users.ts", "x")
	aliases := map[string]string{"@/": "src/"}
	got := resolvePathAlias(dir, "@/api/users", aliases)
	if got == "" {
		t.Fatalf("expected resolution, got empty")
	}
}

func TestResolvePathAlias_PrefixMismatch(t *testing.T) {
	dir := t.TempDir()
	aliases := map[string]string{"@/": "src/"}
	if got := resolvePathAlias(dir, "other/users", aliases); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestResolvePathAlias_FileMissing(t *testing.T) {
	dir := t.TempDir()
	aliases := map[string]string{"@/": "src/"}
	if got := resolvePathAlias(dir, "@/api/missing", aliases); got != "" {
		t.Fatalf("got %q", got)
	}
}
