//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractFuncNameFromCall_Identifier 테스트
package hono

import "testing"

func TestExtractFuncNameFromCall_Identifier(t *testing.T) {
	fi := mustParse(t, []byte("doThing();\n"))
	calls := findAllByType(fi.Root, "call_expression")
	if got := extractFuncNameFromCall(calls[0], fi.Src); got != "doThing" {
		t.Fatalf("got %q", got)
	}
}
