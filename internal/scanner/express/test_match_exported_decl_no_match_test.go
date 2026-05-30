//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchExportedDecl_NoMatch 테스트
package express

import "testing"

func TestMatchExportedDecl_NoMatch(t *testing.T) {
	fi := mustParse(t, []byte(`export function h() {}`))
	if body := matchExportedDecl(topChild(t, fi, "export_statement"), fi.Src, "other"); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}
