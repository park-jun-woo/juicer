//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestAllJSONInsideBlocks_Inside_Round5 테스트
package supafunc

import "testing"

func TestAllJSONInsideBlocks_Inside_Round5(t *testing.T) {

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
	if !allJSONInsideBlocks(body, fi.Src, blocks) {
		t.Fatal("expected true when req.json() inside block")
	}
}
