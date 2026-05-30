//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchAnyDecl_Exported 테스트
package express

import "testing"

func TestMatchAnyDecl_Exported(t *testing.T) {
	fi := mustParse(t, []byte(`export function handler(req, res) { res.json({}); }`))
	if body := matchAnyDecl(topChild(t, fi, "export_statement"), fi.Src, "handler"); body == nil {
		t.Fatal("expected exported body")
	}
}
