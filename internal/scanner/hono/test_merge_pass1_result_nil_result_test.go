//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestMergePass1Result_NilResult 테스트
package hono

import "testing"

func TestMergePass1Result_NilResult(t *testing.T) {

	parsed, vars, bp, schemas, groups, imports := newMergeMaps()
	mergePass1Result("/no/such/file.ts", "/no/such", parsed, vars, bp, schemas, groups, imports)
	if len(parsed) != 0 || len(vars) != 0 || len(bp) != 0 || len(*groups) != 0 {
		t.Fatal("expected no mutation for nil result")
	}
}
