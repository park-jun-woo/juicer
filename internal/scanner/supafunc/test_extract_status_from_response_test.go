//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestExtractStatusFromResponse 테스트
package supafunc

import "testing"

func TestExtractStatusFromResponse(t *testing.T) {
	fi := mustParse(t, []byte(`const r = new Response(JSON.stringify(x), { status: 404 });`))
	news := findAllByType(fi.Root, "new_expression")
	if len(news) == 0 {
		t.Fatal("no new_expression")
	}
	if got := extractStatusFromResponse(news[0], fi.Src); got != "404" {
		t.Fatalf("got %q", got)
	}
}
