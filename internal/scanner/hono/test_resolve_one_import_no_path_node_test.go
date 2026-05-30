//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestResolveOneImport_NoPathNode 테스트
package hono

import "testing"

func TestResolveOneImport_NoPathNode(t *testing.T) {
	dir := t.TempDir()

	imp := resolveOne(t, dir, `const x = 1`)
	if len(imp) != 0 {
		t.Fatalf("expected none, got %v", imp)
	}
}
