//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what TestRemoveColumn 테스트
package ddl

import (
	"testing"
)

func TestRemoveColumn(t *testing.T) {
	cols := []Column{
		{Name: "id", Raw: "id BIGINT"},
		{Name: "name", Raw: "name TEXT"},
		{Name: "email", Raw: "email TEXT"},
	}

	result := removeColumn(cols, "name")
	if len(result) != 2 {
		t.Errorf("expected 2 columns, got %d", len(result))
	}
	for _, c := range result {
		if c.Name == "name" {
			t.Error("column 'name' should have been removed")
		}
	}

	// Removing non-existent column
	result2 := removeColumn(cols, "nonexistent")
	if len(result2) != 3 {
		t.Errorf("expected 3 columns, got %d", len(result2))
	}
}
