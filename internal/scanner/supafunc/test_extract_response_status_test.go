//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestExtractResponseStatus new Response(..., { status: 201 }) 추출 테스트
package supafunc

import (
	"testing"
)

func TestExtractResponseStatus(t *testing.T) {
	src := []byte(`
serve(async (req) => {
  return new Response(JSON.stringify({ ok: true }), {
    headers: { "Content-Type": "application/json" },
    status: 201,
  })
})
`)
	fi := mustParse(t, src)
	body, _ := findServeCallback(fi)
	if body == nil {
		t.Fatal("no callback body found")
	}
	statuses := extractResponseStatus(body, fi.Src)
	if len(statuses) != 1 {
		t.Fatalf("expected 1 status, got %d: %v", len(statuses), statuses)
	}
	if statuses[0] != "201" {
		t.Fatalf("expected '201', got %q", statuses[0])
	}
}
