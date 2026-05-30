//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolvePathAlias_Match 테스트
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
