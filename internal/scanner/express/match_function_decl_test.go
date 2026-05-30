//ff:func feature=scan type=test control=sequence topic=express
//ff:what matchFunctionDecl: 매칭 / 비함수선언 / 이름불일치
package express

import "testing"

func TestMatchFunctionDecl_Match(t *testing.T) {
	fi := mustParse(t, []byte(`function h(req, res) { res.json({}); }`))
	if body := matchFunctionDecl(topChild(t, fi, "function_declaration"), fi.Src, "h"); body == nil {
		t.Fatal("expected body")
	}
}

func TestMatchFunctionDecl_NotFunction(t *testing.T) {
	fi := mustParse(t, []byte(`const h = 1;`))
	if body := matchFunctionDecl(topChild(t, fi, "lexical_declaration"), fi.Src, "h"); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}

func TestMatchFunctionDecl_NameMismatch(t *testing.T) {
	fi := mustParse(t, []byte(`function h() {}`))
	if body := matchFunctionDecl(topChild(t, fi, "function_declaration"), fi.Src, "other"); body != nil {
		t.Fatalf("expected nil, got %v", body.Type())
	}
}
