//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchExportedDecl_NotExport 테스트
package express

import "testing"

func TestMatchExportedDecl_NotExport(t *testing.T) {
	fi := mustParse(t, []byte(`function h() {}`))
	if body := matchExportedDecl(topChild(t, fi, "function_declaration"), fi.Src, "h"); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}
