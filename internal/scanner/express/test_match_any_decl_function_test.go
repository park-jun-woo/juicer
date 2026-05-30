//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchAnyDecl_Function 테스트
package express

import "testing"

func TestMatchAnyDecl_Function(t *testing.T) {
	fi := mustParse(t, []byte(`function handler(req, res) { res.json({}); }`))
	if body := matchAnyDecl(topChild(t, fi, "function_declaration"), fi.Src, "handler"); body == nil {
		t.Fatal("expected function body")
	}
}
