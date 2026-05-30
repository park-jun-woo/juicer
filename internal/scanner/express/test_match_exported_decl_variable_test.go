//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchExportedDecl_Variable 테스트
package express

import "testing"

func TestMatchExportedDecl_Variable(t *testing.T) {
	fi := mustParse(t, []byte(`export const h = (req, res) => { res.json({}); };`))
	if body := matchExportedDecl(topChild(t, fi, "export_statement"), fi.Src, "h"); body == nil {
		t.Fatal("expected body")
	}
}
