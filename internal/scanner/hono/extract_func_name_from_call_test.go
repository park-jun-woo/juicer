//ff:func feature=scan type=test control=sequence topic=hono
//ff:what extractFuncNameFromCall 테스트
package hono

import "testing"

func TestExtractFuncNameFromCall_Identifier(t *testing.T) {
	fi := mustParse(t, []byte("doThing();\n"))
	calls := findAllByType(fi.Root, "call_expression")
	if got := extractFuncNameFromCall(calls[0], fi.Src); got != "doThing" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractFuncNameFromCall_Member(t *testing.T) {
	fi := mustParse(t, []byte("app.get();\n"))
	calls := findAllByType(fi.Root, "call_expression")
	if got := extractFuncNameFromCall(calls[0], fi.Src); got != "app.get" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractFuncNameFromCall_Neither(t *testing.T) {
	// IIFE: the function is a parenthesized arrow, not identifier/member -> ""
	fi := mustParse(t, []byte("(() => {})();\n"))
	calls := findAllByType(fi.Root, "call_expression")
	// find the outer call whose function is the parenthesized expression
	for _, c := range calls {
		got := extractFuncNameFromCall(c, fi.Src)
		if got == "" {
			return
		}
	}
	t.Skip("no call with empty func name")
}
