//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolvePathAlias_PrefixMismatch 테스트
package express

import "testing"

func TestResolvePathAlias_PrefixMismatch(t *testing.T) {
	dir := t.TempDir()
	aliases := map[string]string{"@/": "src/"}
	if got := resolvePathAlias(dir, "other/users", aliases); got != "" {
		t.Fatalf("got %q", got)
	}
}
