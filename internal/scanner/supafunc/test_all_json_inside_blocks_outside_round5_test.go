//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestAllJSONInsideBlocks_Outside_Round5 테스트
package supafunc

import "testing"

func TestAllJSONInsideBlocks_Outside_Round5(t *testing.T) {

	src := []byte(`
serve(async (req) => {
  const { name } = await req.json()
  if (req.method === "POST") {
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
	if allJSONInsideBlocks(body, fi.Src, blocks) {
		t.Fatal("expected false when req.json() outside blocks")
	}
}
