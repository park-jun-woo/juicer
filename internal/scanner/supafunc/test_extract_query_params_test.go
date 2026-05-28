//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestExtractQueryParams searchParams.get("page") 추출 테스트
package supafunc

import (
	"testing"
)

func TestExtractQueryParams(t *testing.T) {
	src := []byte(`
serve(async (req) => {
  const url = new URL(req.url)
  const page = url.searchParams.get("page")
  const limit = url.searchParams.get("limit")
  return new Response(JSON.stringify({ page, limit }))
})
`)
	fi := mustParse(t, src)
	body, _ := findServeCallback(fi)
	if body == nil {
		t.Fatal("no callback body found")
	}
	params := extractQueryParams(body, fi.Src)
	if len(params) != 2 {
		t.Fatalf("expected 2 params, got %d: %v", len(params), params)
	}
	if params[0] != "page" {
		t.Fatalf("expected 'page', got %q", params[0])
	}
	if params[1] != "limit" {
		t.Fatalf("expected 'limit', got %q", params[1])
	}
}
