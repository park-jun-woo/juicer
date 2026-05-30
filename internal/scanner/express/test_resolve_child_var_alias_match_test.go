//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveChildVar_AliasMatch 테스트
package express

import "testing"

func TestResolveChildVar_AliasMatch(t *testing.T) {
	all := map[string]map[string]bool{"c.ts": {"users": true, "admin": true}}
	if got := resolveChildVar("c.ts", "users", all); got != "users" {
		t.Fatalf("got %q", got)
	}
}
