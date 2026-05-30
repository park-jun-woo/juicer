//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveChildVar_None 테스트
package express

import "testing"

func TestResolveChildVar_None(t *testing.T) {
	if got := resolveChildVar("c.ts", "x", map[string]map[string]bool{}); got != "" {
		t.Fatalf("got %q", got)
	}
}
