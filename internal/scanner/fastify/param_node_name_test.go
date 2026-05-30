//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what paramNodeName 테스트
package fastify

import "testing"

func TestParamNodeName_Identifier(t *testing.T) {
	// A bare identifier node (e.g. from extractFirstParamName iterating an
	// arrow with a single bare param). Construct one directly from an
	// identifier reference.
	fi := mustParse(t, []byte("foo(myParam);\n"))
	args := findAllByType(fi.Root, "arguments")[0]
	id := findChildByType(args, "identifier")
	if id == nil {
		t.Fatal("no identifier")
	}
	if got := paramNodeName(id, fi.Src); got != "myParam" {
		t.Fatalf("identifier: got %q", got)
	}
}

func TestParamNodeName_RequiredParameter(t *testing.T) {
	fi := mustParse(t, []byte("function f(a: number) {}\n"))
	params := findAllByType(fi.Root, "formal_parameters")[0]
	rp := findChildByType(params, "required_parameter")
	if rp == nil {
		t.Skip("no required_parameter")
	}
	if got := paramNodeName(rp, fi.Src); got != "a" {
		t.Fatalf("required_parameter: got %q", got)
	}
}

func TestParamNodeName_OptionalParameter(t *testing.T) {
	fi := mustParse(t, []byte("function f(a?: number) {}\n"))
	params := findAllByType(fi.Root, "formal_parameters")[0]
	op := findChildByType(params, "optional_parameter")
	if op == nil {
		t.Skip("no optional_parameter")
	}
	if got := paramNodeName(op, fi.Src); got != "a" {
		t.Fatalf("optional_parameter: got %q", got)
	}
}

func TestParamNodeName_Other(t *testing.T) {
	// a punctuation node ("(") yields ""
	fi := mustParse(t, []byte("function f(a) {}\n"))
	params := findAllByType(fi.Root, "formal_parameters")[0]
	if got := paramNodeName(params.Child(0), fi.Src); got != "" {
		t.Fatalf("expected empty for '(' node, got %q", got)
	}
}
