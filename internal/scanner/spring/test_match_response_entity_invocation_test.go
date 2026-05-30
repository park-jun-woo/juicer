//ff:func feature=scan type=test control=iteration dimension=1 topic=spring
//ff:what TestMatchResponseEntityInvocation 테스트
package spring

import "testing"

func TestMatchResponseEntityInvocation(t *testing.T) {
	root, src := parseS(t, `class C { void m() { return ResponseEntity.created(uri).build(); } }`)
	invs := findAllByType(root, "method_invocation")
	for _, inv := range invs {
		if matchResponseEntityInvocation(inv, src) == "201" {
			return
		}
	}
	t.Fatal("did not match 201")
}
