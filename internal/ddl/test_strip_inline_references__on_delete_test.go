//ff:func feature=ddl type=test control=sequence
//ff:what stripInlineReferences가 ON DELETE CASCADE까지 제거하는지 테스트
package ddl

import "testing"

func TestStripInlineReferences_OnDelete(t *testing.T) {
	raw := "tenant_id BIGINT NOT NULL REFERENCES tenants(id) ON DELETE CASCADE"
	got := stripInlineReferences(raw)
	want := "tenant_id BIGINT NOT NULL"
	if got != want {
		t.Fatalf("stripInlineReferences = %q, want %q", got, want)
	}
}
