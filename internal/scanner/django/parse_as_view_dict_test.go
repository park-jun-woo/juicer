//ff:func feature=scan type=test control=sequence topic=django
//ff:what parseAsViewDict — as_view({method:action}) dict 파싱을 검증
package django

import "testing"

func TestParseAsViewDict(t *testing.T) {
	src := []byte("path('x/', V.as_view({\"get\": \"list\", \"POST\": \"create\"}))\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	// The .as_view(...) call is the inner call node.
	args := djFirst(t, root, "argument_list")
	pos := positionalArgs(args)
	got := parseAsViewDict(pos[1], src)
	if len(got) != 2 {
		t.Fatalf("expected 2 entries, got %v", got)
	}
	if got["get"] != "list" {
		t.Errorf("get -> %q, want list", got["get"])
	}
	if got["post"] != "create" {
		t.Errorf("post (lowercased) -> %q, want create", got["post"])
	}

	// No-dict as_view returns nil.
	src2 := []byte("path('x/', V.as_view())\n")
	root2, _ := parsePython(src2)
	args2 := djFirst(t, root2, "argument_list")
	pos2 := positionalArgs(args2)
	if got := parseAsViewDict(pos2[1], src2); got != nil {
		t.Errorf("expected nil for empty as_view, got %v", got)
	}
}
