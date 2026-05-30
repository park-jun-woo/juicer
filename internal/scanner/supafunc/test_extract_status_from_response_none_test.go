//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestExtractStatusFromResponse_None 테스트
package supafunc

import "testing"

func TestExtractStatusFromResponse_None(t *testing.T) {
	fi := mustParse(t, []byte(`const r = new Response("ok");`))
	news := findAllByType(fi.Root, "new_expression")
	if got := extractStatusFromResponse(news[0], fi.Src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
