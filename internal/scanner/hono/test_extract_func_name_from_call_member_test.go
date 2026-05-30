//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractFuncNameFromCall_Member 테스트
package hono

import "testing"

func TestExtractFuncNameFromCall_Member(t *testing.T) {
	fi := mustParse(t, []byte("app.get();\n"))
	calls := findAllByType(fi.Root, "call_expression")
	if got := extractFuncNameFromCall(calls[0], fi.Src); got != "app.get" {
		t.Fatalf("got %q", got)
	}
}
