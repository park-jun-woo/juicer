//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestExtractMethodBlockFromIf_Round5 테스트
package supafunc

import "testing"

func TestExtractMethodBlockFromIf_Round5(t *testing.T) {
	src := []byte(`
serve(async (req) => {
  if (req.method === "POST") {
    const { name } = await req.json()
    return new Response("ok", { status: 201 })
  }
})
`)
	fi := mustParse(t, src)
	body, _ := findServeCallback(fi)
	if body == nil {
		t.Fatal("no callback body")
	}
	blocks := extractMethodBlock(body, fi.Src)
	if blocks["POST"] == nil {
		t.Fatalf("expected POST block, got %v keys", blocks)
	}
}
