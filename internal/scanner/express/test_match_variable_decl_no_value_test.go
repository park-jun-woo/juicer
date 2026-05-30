//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchVariableDecl_NoValue 테스트
package express

import "testing"

func TestMatchVariableDecl_NoValue(t *testing.T) {

	fi := mustParse(t, []byte(`let h;`))
	if body := matchVariableDecl(topChild(t, fi, "lexical_declaration"), fi.Src, "h"); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}
