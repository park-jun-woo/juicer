//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractScopes_None 테스트
package actix

import "testing"

func TestExtractScopes_None(t *testing.T) {
	src := []byte(`fn f() { web::resource("/x"); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fi := &fileInfo{src: src, root: root}
	if scopes := extractScopes(fi); len(scopes) != 0 {
		t.Fatalf("expected no scopes, got %+v", scopes)
	}
}
