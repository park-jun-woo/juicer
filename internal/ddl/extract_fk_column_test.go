//ff:func feature=ddl type=test control=sequence
//ff:what extractFKColumn의 단일 컬럼 FK 추출 테스트
package ddl

import "testing"

func TestExtractFKColumn(t *testing.T) {
	got := extractFKColumn("CONSTRAINT building_owners_owner_id_fkey FOREIGN KEY (owner_id) REFERENCES users(id)")
	if got != "owner_id" {
		t.Fatalf("extractFKColumn = %q, want %q", got, "owner_id")
	}
}
