//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractResourcePath 테스트
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
