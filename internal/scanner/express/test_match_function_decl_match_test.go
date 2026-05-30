//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchFunctionDecl_Match 테스트
package express

import "testing"

func TestMatchFunctionDecl_Match(t *testing.T) {
	fi := mustParse(t, []byte(`function h(req, res) { res.json({}); }`))
	if body := matchFunctionDecl(topChild(t, fi, "function_declaration"), fi.Src, "h"); body == nil {
		t.Fatal("expected body")
	}
}
