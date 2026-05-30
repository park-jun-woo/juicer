//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestUnwrapNewExpression_NewRouterIdent 테스트
package express

import "testing"

func TestUnwrapNewExpression_NewRouterIdent(t *testing.T) {

	fi := mustParse(t, []byte(`const r = new Router();`))
	ne := findAllByType(fi.Root, "new_expression")[0]
	got := unwrapNewExpression(ne, fi.Src, map[string]bool{"Router": true})
	if got == nil || got != ne {
		t.Fatalf("expected node itself, got %v", got)
	}
}
