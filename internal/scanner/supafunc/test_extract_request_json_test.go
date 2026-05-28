//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestExtractRequestJSON const { id, name } = await req.json() 필드 추출 테스트
package supafunc

import (
	"testing"
)

func TestExtractRequestJSON(t *testing.T) {
	src := []byte(`
serve(async (req) => {
  const { id, name } = await req.json()
  return new Response(JSON.stringify({ id, name }))
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
	if fields[0] != "id" {
		t.Fatalf("expected first field 'id', got %q", fields[0])
	}
	if fields[1] != "name" {
		t.Fatalf("expected second field 'name', got %q", fields[1])
	}
}
