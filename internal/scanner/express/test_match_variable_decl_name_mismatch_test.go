//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchVariableDecl_NameMismatch 테스트
package express

import "testing"

func TestMatchVariableDecl_NameMismatch(t *testing.T) {
	fi := mustParse(t, []byte(`const h = () => {};`))
	if body := matchVariableDecl(topChild(t, fi, "lexical_declaration"), fi.Src, "other"); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}
