//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what splitTopLevel 테스트
package fastapi

import "testing"

func TestSplitTopLevel(t *testing.T) {
	got := splitTopLevel("str, int, None")
	if len(got) != 3 || got[0] != "str" || got[1] != "int" || got[2] != "None" {
		t.Fatalf("got %v", got)
	}

	// nested brackets
	got2 := splitTopLevel("List[str], None")
	if len(got2) != 2 || got2[0] != "List[str]" {
		t.Fatalf("got %v", got2)
	}

	// pipe
	got3 := splitTopLevel("str | None")
	if len(got3) != 2 || got3[0] != "str" || got3[1] != "None" {
		t.Fatalf("got %v", got3)
	}

	// empty
	got4 := splitTopLevel("")
	if len(got4) != 0 {
		t.Fatalf("expected empty, got %v", got4)
	}
}
