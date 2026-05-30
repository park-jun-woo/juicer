//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestUnwrapNewExpression_NoMatch 테스트
package express

import "testing"

func TestUnwrapNewExpression_NoMatch(t *testing.T) {
	fi := mustParse(t, []byte(`const r = new Other;`))
	ne := findAllByType(fi.Root, "new_expression")
	if len(ne) == 0 {
		t.Skip("no new_expression")
	}
	if got := unwrapNewExpression(ne[0], fi.Src, map[string]bool{"Router": true}); got != nil {
		t.Fatalf("expected nil, got %v", got.Type())
	}
}
