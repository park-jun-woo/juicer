//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestResolveOneImport_External 테스트
package hono

import "testing"

func TestResolveOneImport_External(t *testing.T) {
	dir := t.TempDir()
	imp := resolveOne(t, dir, `import { z } from "zod"`)
	if len(imp) != 0 {
		t.Fatalf("expected external skipped, got %v", imp)
	}
}
