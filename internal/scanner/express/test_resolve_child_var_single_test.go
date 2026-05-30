//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveChildVar_Single 테스트
package express

import "testing"

func TestResolveChildVar_Single(t *testing.T) {
	all := map[string]map[string]bool{"c.ts": {"router": true}}
	if got := resolveChildVar("c.ts", "x", all); got != "router" {
		t.Fatalf("got %q", got)
	}
}
