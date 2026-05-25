package sqls

import "testing"

func TestAppendUnique_New(t *testing.T) {
	result := appendUnique([]string{"a"}, "b")
	if len(result) != 2 {
		t.Fatal("expected 2")
	}
}

func TestAppendUnique_Duplicate(t *testing.T) {
	result := appendUnique([]string{"a"}, "a")
	if len(result) != 1 {
		t.Fatal("expected 1")
	}
}
