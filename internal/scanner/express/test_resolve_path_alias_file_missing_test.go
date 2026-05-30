//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolvePathAlias_FileMissing 테스트
package express

import "testing"

func TestResolvePathAlias_FileMissing(t *testing.T) {
	dir := t.TempDir()
	aliases := map[string]string{"@/": "src/"}
	if got := resolvePathAlias(dir, "@/api/missing", aliases); got != "" {
		t.Fatalf("got %q", got)
	}
}
