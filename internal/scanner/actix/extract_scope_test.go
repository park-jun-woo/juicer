//ff:func feature=scan type=test control=sequence topic=actix
//ff:what extractScopes — web::scope prefix/핸들러 추출을 검증
package actix

import "testing"

func TestExtractScopes(t *testing.T) {
	src := []byte(`
fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(
        web::scope("/api")
            .service(list_users)
            .service(create_user)
    );
}
`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fi := &fileInfo{src: src, root: root}
	scopes := extractScopes(fi)
	if len(scopes) != 1 {
		t.Fatalf("expected 1 scope, got %d: %+v", len(scopes), scopes)
	}
	if scopes[0].prefix != "/api" {
		t.Errorf("prefix = %q, want /api", scopes[0].prefix)
	}
	if len(scopes[0].handlers) != 2 {
		t.Errorf("handlers = %v, want 2", scopes[0].handlers)
	}
}

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
