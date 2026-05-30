//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchFunctionDecl_NotFunction 테스트
package express

import "testing"

func TestMatchFunctionDecl_NotFunction(t *testing.T) {
	fi := mustParse(t, []byte(`const h = 1;`))
	if body := matchFunctionDecl(topChild(t, fi, "lexical_declaration"), fi.Src, "h"); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}
