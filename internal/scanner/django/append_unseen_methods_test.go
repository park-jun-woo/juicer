//ff:func feature=scan type=test control=sequence topic=django
//ff:what appendUnseenMethods — 중복 제거 후 actionMethod 추가를 검증
package django

import "testing"

func TestAppendUnseenMethods(t *testing.T) {
	seen := map[string]bool{}
	ms := []actionMethod{
		{action: "list", method: "GET"},
		{action: "create", method: "POST"},
		{action: "list", method: "GET"}, // duplicate -> skipped
	}
	out := appendUnseenMethods(nil, ms, seen)
	if len(out) != 2 {
		t.Fatalf("expected 2 unique methods, got %d: %+v", len(out), out)
	}

	// Calling again with an already-seen method adds nothing.
	out = appendUnseenMethods(out, []actionMethod{{action: "list", method: "GET"}}, seen)
	if len(out) != 2 {
		t.Fatalf("expected still 2, got %d", len(out))
	}
}
