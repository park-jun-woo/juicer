//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchAnyDecl_Variable 테스트
package express

import "testing"

func TestMatchAnyDecl_Variable(t *testing.T) {
	fi := mustParse(t, []byte(`const handler = (req, res) => { res.json({}); };`))
	if body := matchAnyDecl(topChild(t, fi, "lexical_declaration"), fi.Src, "handler"); body == nil {
		t.Fatal("expected variable body")
	}
}
