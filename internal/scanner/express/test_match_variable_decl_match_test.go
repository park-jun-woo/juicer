//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchVariableDecl_Match 테스트
package express

import "testing"

func TestMatchVariableDecl_Match(t *testing.T) {
	fi := mustParse(t, []byte(`const h = (req, res) => { res.json({}); };`))
	if body := matchVariableDecl(topChild(t, fi, "lexical_declaration"), fi.Src, "h"); body == nil {
		t.Fatal("expected body")
	}
}
