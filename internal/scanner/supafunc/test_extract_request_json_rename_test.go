//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestExtractRequestJSON_PairPattern pair_pattern 프로퍼티 리네이밍에서 원본 필드명 추출 테스트
package supafunc

import (
	"testing"
)

func TestExtractRequestJSON_PairPattern(t *testing.T) {
	src := []byte(`
serve(async (req) => {
  const { store: storeId, items } = await req.json()
  return new Response(JSON.stringify({ storeId, items }))
})
`)
	fi := mustParse(t, src)
	body, _ := findServeCallback(fi)
	if body == nil {
		t.Fatal("no callback body found")
	}
	fields := extractRequestJSON(body, fi.Src)
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d: %v", len(fields), fields)
	}
	if fields[0] != "store" {
		t.Fatalf("expected first field 'store', got %q", fields[0])
	}
	if fields[1] != "items" {
		t.Fatalf("expected second field 'items', got %q", fields[1])
	}
}
