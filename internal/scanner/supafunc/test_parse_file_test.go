//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestParseFile 테스트
package supafunc

import "testing"

func TestParseFile(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "index.ts", `serve(async (req) => new Response("ok"));`)
	fi, err := parseFile(dir + "/index.ts")
	if err != nil || fi == nil || fi.Root == nil {
		t.Fatalf("err: %v", err)
	}
}
