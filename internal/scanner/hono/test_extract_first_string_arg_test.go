//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractFirstStringArg 테스트
package hono

import "testing"

func TestExtractFirstStringArg(t *testing.T) {
	fi := mustParse(t, []byte(`f("/api", x);`+"\n"))
	calls := findAllByType(fi.Root, "call_expression")
	if got := extractFirstStringArg(calls[0], fi.Src); got != "/api" {
		t.Fatalf("got %q", got)
	}
}
