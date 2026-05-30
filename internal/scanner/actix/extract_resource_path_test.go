//ff:func feature=scan type=test control=sequence topic=actix
//ff:what extractResourcePath — web::resource("...") 경로 추출을 검증
package actix

import "testing"

func TestExtractResourcePath(t *testing.T) {
	src := []byte(`
fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(web::resource("/widgets/{id}").route(web::get().to(get_widget)));
}
`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	if got := extractResourcePath(root, src); got != "/widgets/{id}" {
		t.Fatalf("extractResourcePath = %q, want /widgets/{id}", got)
	}
}

func TestExtractResourcePath_None(t *testing.T) {
	src := []byte(`fn f() { web::scope("/p"); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	if got := extractResourcePath(root, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
