//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchFunctionDecl_NameMismatch 테스트
package express

import "testing"

func TestMatchFunctionDecl_NameMismatch(t *testing.T) {
	fi := mustParse(t, []byte(`function h() {}`))
	if body := matchFunctionDecl(topChild(t, fi, "function_declaration"), fi.Src, "other"); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}
