//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestExtractBodyMemberAccess_Round5 테스트
package supafunc

import "testing"

func TestExtractBodyMemberAccess_Round5(t *testing.T) {
	src := []byte(`
serve(async (req) => {
  const body = await req.json()
  const name = body.name
  const age = body.age
  return new Response("ok", { status: 200 })
})
`)
	fi := mustParse(t, src)
	body, _ := findServeCallback(fi)
	if body == nil {
		t.Fatal("no callback body")
	}
	fields := extractBodyMemberAccess(body, fi.Src)
	if len(fields) == 0 {
		t.Fatalf("expected member-access fields, got %v", fields)
	}
}
