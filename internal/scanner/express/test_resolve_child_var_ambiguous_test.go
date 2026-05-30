//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveChildVar_Ambiguous 테스트
package express

import "testing"

func TestResolveChildVar_Ambiguous(t *testing.T) {
	all := map[string]map[string]bool{"c.ts": {"users": true, "admin": true}}
	if got := resolveChildVar("c.ts", "nope", all); got != "" {
		t.Fatalf("got %q", got)
	}
}
