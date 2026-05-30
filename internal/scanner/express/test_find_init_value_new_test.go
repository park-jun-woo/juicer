//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestFindInitValue_New 테스트
package express

import "testing"

func TestFindInitValue_New(t *testing.T) {
	fi := mustParse(t, []byte(`const r = new Router();`))
	if v := findInitValue(firstDeclarator(t, fi)); v == nil || v.Type() != "new_expression" {
		t.Fatalf("got %v", v)
	}
}
