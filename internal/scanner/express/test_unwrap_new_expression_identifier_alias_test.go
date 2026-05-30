//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestUnwrapNewExpression_IdentifierAlias 테스트
package express

import "testing"

func TestUnwrapNewExpression_IdentifierAlias(t *testing.T) {
	fi := mustParse(t, []byte(`const r = new Router;`))
	ne := findAllByType(fi.Root, "new_expression")
	if len(ne) == 0 {
		t.Skip("no new_expression for bare new")
	}
	got := unwrapNewExpression(ne[0], fi.Src, map[string]bool{"Router": true})
	if got == nil {
		t.Fatal("expected non-nil for identifier alias")
	}
}
