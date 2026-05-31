//ff:func feature=scan type=test topic=actix control=iteration dimension=1
//ff:what applyResponseBody .json(Struct{...}) 응답 본문 타입/필드 반영 테스트
package actix

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestApplyResponseBody(t *testing.T) {
	src := []byte(`fn handler() -> HttpResponse {
    HttpResponse::Ok().json(UserResponse { id: 1 })
}`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	block := findAllByType(root, "block")[0]
	scoped := findAllByType(root, "scoped_identifier")
	if len(scoped) == 0 {
		t.Fatal("no scoped_identifier")
	}
	ctx := &responseCtx{block: block, src: src, cache: map[string][]scanner.Field{}}

	var resp scanner.Response
	// find the scoped identifier that drives the .json(...) chain (HttpResponse::Ok)
	for _, sid := range scoped {
		applyResponseBody(&resp, sid, ctx)
		if resp.TypeName != "" {
			break
		}
	}
	if resp.TypeName != "UserResponse" {
		t.Errorf("TypeName = %q, want UserResponse", resp.TypeName)
	}
}
