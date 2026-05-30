//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractCallStringArg_String 테스트
package fastify

import "testing"

func TestExtractCallStringArg_String(t *testing.T) {
	fi := mustParse(t, []byte(`require("@fastify/autoload");`+"\n"))
	calls := findAllByType(fi.Root, "call_expression")
	if len(calls) == 0 {
		t.Fatal("no call")
	}
	if got := extractCallStringArg(calls[0], fi.Src); got != "@fastify/autoload" {
		t.Fatalf("got %q", got)
	}
}
