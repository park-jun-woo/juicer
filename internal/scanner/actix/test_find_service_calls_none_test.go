//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestFindServiceCalls_None 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

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
