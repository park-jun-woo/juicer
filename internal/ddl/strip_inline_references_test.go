//ff:func feature=ddl type=test control=sequence
//ff:what stripInlineReferences의 기본 REFERENCES 제거 테스트
package ddl

import "testing"

func TestStripInlineReferences(t *testing.T) {
	raw := "owner_id BIGINT NOT NULL REFERENCES owners(id)"
	got := stripInlineReferences(raw)
	want := "owner_id BIGINT NOT NULL"
	if got != want {
		t.Fatalf("stripInlineReferences = %q, want %q", got, want)
	}
}
