//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractResponses_WithBlock 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestExtractResponses_WithBlock(t *testing.T) {
	src := []byte(`
struct CreatedURL { a: String }
async fn create_url() -> HttpResponse {
    HttpResponse::Created().json(CreatedURL{a: x})
}
`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fn := findFuncByName(root, src, "create_url")
	if fn == nil {
		t.Fatal("function not found")
	}
	idx := buildStructIndex([]*fileInfo{{src: src, root: root}})
	cache := map[string][]scanner.Field{}
	resps := extractResponses(fn, src, idx, cache)
	if len(resps) != 1 {
		t.Fatalf("expected 1 response, got %d: %+v", len(resps), resps)
	}
	if resps[0].Status != "201" {
		t.Errorf("status = %q, want 201", resps[0].Status)
	}
}
