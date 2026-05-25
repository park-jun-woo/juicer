//ff:func feature=ddl type=extract control=sequence
//ff:what TestRemoveColumn_Empty 테스트
package ddl

import "testing"

func TestRemoveColumn_Empty(t *testing.T) {
	result := removeColumn(nil, "id")
	if len(result) != 0 {
		t.Fatalf("expected 0, got %d", len(result))
	}
}
