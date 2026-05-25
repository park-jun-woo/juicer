//ff:func feature=ddl type=extract control=sequence
//ff:what TestRemoveColumn_NotFound 테스트
package ddl

import "testing"

func TestRemoveColumn_NotFound(t *testing.T) {
	cols := []Column{{Name: "id", Raw: "id INT"}}
	result := removeColumn(cols, "email")
	if len(result) != 1 {
		t.Fatalf("expected 1, got %d", len(result))
	}
}
