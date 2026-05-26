//ff:func feature=ddl type=test control=sequence
//ff:what extractFKColumn에서 FK가 없는 UNIQUE 제약 시 빈 문자열 반환 테스트
package ddl

import "testing"

func TestExtractFKColumn_NoFK(t *testing.T) {
	got := extractFKColumn("CONSTRAINT uq_billing UNIQUE (contract_id, billing_month)")
	if got != "" {
		t.Fatalf("extractFKColumn = %q, want empty", got)
	}
}
