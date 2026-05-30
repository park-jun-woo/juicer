//ff:func feature=scan type=test control=sequence topic=express
//ff:what resolveChildVar: 없음 / 단일 / alias매칭 / 모호 분기
package express

import "testing"

func TestResolveChildVar_None(t *testing.T) {
	if got := resolveChildVar("c.ts", "x", map[string]map[string]bool{}); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestResolveChildVar_Single(t *testing.T) {
	all := map[string]map[string]bool{"c.ts": {"router": true}}
	if got := resolveChildVar("c.ts", "x", all); got != "router" {
		t.Fatalf("got %q", got)
	}
}

func TestResolveChildVar_AliasMatch(t *testing.T) {
	all := map[string]map[string]bool{"c.ts": {"users": true, "admin": true}}
	if got := resolveChildVar("c.ts", "users", all); got != "users" {
		t.Fatalf("got %q", got)
	}
}

func TestResolveChildVar_Ambiguous(t *testing.T) {
	all := map[string]map[string]bool{"c.ts": {"users": true, "admin": true}}
	if got := resolveChildVar("c.ts", "nope", all); got != "" {
		t.Fatalf("got %q", got)
	}
}
