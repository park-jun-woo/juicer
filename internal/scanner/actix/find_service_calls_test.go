//ff:func feature=scan type=test control=sequence topic=actix
//ff:what findServiceCalls — .service() 호출 인자 콜백 전달을 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestFindServiceCalls(t *testing.T) {
	src := []byte(`
fn f() {
    web::scope("/x")
        .service(a)
        .service(b);
}
`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	chain := findCallByFuncSuffix(root, src, ".service")
	if chain == nil {
		t.Fatal("no .service call found")
	}
	count := 0
	findServiceCalls(chain, src, func(arg *sitter.Node) { count++ })
	if count != 2 {
		t.Fatalf("expected 2 service calls, got %d", count)
	}
}

func TestFindServiceCalls_None(t *testing.T) {
	src := []byte(`fn f() { web::scope("/x").route(web::get().to(a)); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	chain := findCallByFuncSuffix(root, src, ".route")
	if chain == nil {
		t.Fatal("no .route call found")
	}
	count := 0
	findServiceCalls(chain, src, func(arg *sitter.Node) { count++ })
	if count != 0 {
		t.Fatalf("expected 0 service calls, got %d", count)
	}
}
