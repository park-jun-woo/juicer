//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestResolveOneImport_Unresolvable 테스트
package hono

import "testing"

func TestResolveOneImport_Unresolvable(t *testing.T) {
	dir := t.TempDir()
	imp := resolveOne(t, dir, `import { x } from "./nope"`)
	if len(imp) != 0 {
		t.Fatalf("expected unresolved skipped, got %v", imp)
	}
}
