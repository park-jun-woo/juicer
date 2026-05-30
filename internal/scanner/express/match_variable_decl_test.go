//ff:func feature=scan type=test control=sequence topic=express
//ff:what matchVariableDecl: 매칭 / 비변수선언 / 이름불일치 / 값없음
package express

import "testing"

func TestMatchVariableDecl_Match(t *testing.T) {
	fi := mustParse(t, []byte(`const h = (req, res) => { res.json({}); };`))
	if body := matchVariableDecl(topChild(t, fi, "lexical_declaration"), fi.Src, "h"); body == nil {
		t.Fatal("expected body")
	}
}

func TestMatchVariableDecl_NotVarDecl(t *testing.T) {
	fi := mustParse(t, []byte(`function h() {}`))
	if body := matchVariableDecl(topChild(t, fi, "function_declaration"), fi.Src, "h"); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}

func TestMatchVariableDecl_NameMismatch(t *testing.T) {
	fi := mustParse(t, []byte(`const h = () => {};`))
	if body := matchVariableDecl(topChild(t, fi, "lexical_declaration"), fi.Src, "other"); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}

func TestMatchVariableDecl_NoValue(t *testing.T) {
	// declarator without initializer -> value nil
	fi := mustParse(t, []byte(`let h;`))
	if body := matchVariableDecl(topChild(t, fi, "lexical_declaration"), fi.Src, "h"); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}
