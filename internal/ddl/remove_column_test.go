//ff:func feature=ddl type=extract control=sequence
//ff:what TestRemoveColumn_Found 테스트
package ddl

import "testing"

func TestRemoveColumn_Found(t *testing.T) {
	cols := []Column{{Name: "id", Raw: "id INT"}, {Name: "name", Raw: "name TEXT"}}
	result := removeColumn(cols, "name")
	if len(result) != 1 {
		t.Fatalf("expected 1, got %d", len(result))
	}
}
