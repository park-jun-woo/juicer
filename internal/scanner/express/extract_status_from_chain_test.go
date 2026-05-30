//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractStatusFromChain: 정상 / 비멤버 / 비체인 / status아님 / 비number 분기
package express

import "testing"

func TestExtractStatusFromChain_Valid(t *testing.T) {
	fi := mustParse(t, []byte(`res.status(201).json({});`))
	if got := extractStatusFromChain(outermostCall(fi), fi.Src); got != "201" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractStatusFromChain_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`foo();`))
	if got := extractStatusFromChain(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractStatusFromChain_ObjectNotCall(t *testing.T) {
	// res.json(...) -> object is identifier, not call_expression
	fi := mustParse(t, []byte(`res.json({});`))
	if got := extractStatusFromChain(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractStatusFromChain_NotStatusCall(t *testing.T) {
	// inner call is foo() not res.status()
	fi := mustParse(t, []byte(`foo().json({});`))
	if got := extractStatusFromChain(outermostCall(fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractStatusFromChain_NotNumberArg(t *testing.T) {
	fi := mustParse(t, []byte(`res.status(code).json({});`))
	if got := extractStatusFromChain(outermostCall(fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractStatusFromChain_EmptyArgs(t *testing.T) {
	fi := mustParse(t, []byte(`res.status().json({});`))
	if got := extractStatusFromChain(outermostCall(fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
