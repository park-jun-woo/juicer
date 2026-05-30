//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchExportedDecl_Function 테스트
package express

import "testing"

func TestMatchExportedDecl_Function(t *testing.T) {
	fi := mustParse(t, []byte(`export function h(req, res) { res.json({}); }`))
	if body := matchExportedDecl(topChild(t, fi, "export_statement"), fi.Src, "h"); body == nil {
		t.Fatal("expected body")
	}
}
