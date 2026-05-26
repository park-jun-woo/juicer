//ff:func feature=ddl type=test control=sequence
//ff:what REFERENCES가 없는 컬럼은 stripInlineReferences가 변경하지 않는 테스트
package ddl

import "testing"

func TestStripInlineReferences_NoRef(t *testing.T) {
	raw := "name TEXT NOT NULL"
	got := stripInlineReferences(raw)
	if got != raw {
		t.Fatalf("stripInlineReferences = %q, want %q (unchanged)", got, raw)
	}
}
