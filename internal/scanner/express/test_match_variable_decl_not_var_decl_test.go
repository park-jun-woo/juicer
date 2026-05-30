//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestMatchVariableDecl_NotVarDecl 테스트
package express

import "testing"

func TestMatchVariableDecl_NotVarDecl(t *testing.T) {
	fi := mustParse(t, []byte(`function h() {}`))
	if body := matchVariableDecl(topChild(t, fi, "function_declaration"), fi.Src, "h"); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}
