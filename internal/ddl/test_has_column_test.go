//ff:func feature=ddl type=parse control=sequence
//ff:what TestHasColumn 테스트
package ddl

import (
	"testing"
)

func TestHasColumn(t *testing.T) {
	tbl := &Table{
		Columns: []Column{
			{Name: "id", Raw: "id BIGINT"},
			{Name: "name", Raw: "name TEXT"},
		},
	}

	if !hasColumn(tbl, "id") {
		t.Error("expected hasColumn(id) = true")
	}
	if !hasColumn(tbl, "name") {
		t.Error("expected hasColumn(name) = true")
	}
	if hasColumn(tbl, "email") {
		t.Error("expected hasColumn(email) = false")
	}
	if hasColumn(&Table{}, "anything") {
		t.Error("expected hasColumn on empty table = false")
	}
}
