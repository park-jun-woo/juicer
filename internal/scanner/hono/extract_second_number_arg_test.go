//ff:func feature=scan type=test control=sequence topic=hono
//ff:what extractSecondNumberArg 테스트
package hono

import "testing"

func TestExtractSecondNumberArg_Number(t *testing.T) {
	call, src := firstCallExpr(t, `c.json({}, 201);`)
	if got := extractSecondNumberArg(call, src, "200"); got != "201" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractSecondNumberArg_NotNumber(t *testing.T) {
	call, src := firstCallExpr(t, `c.json({}, opts);`)
	if got := extractSecondNumberArg(call, src, "200"); got != "200" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractSecondNumberArg_OneArg(t *testing.T) {
	call, src := firstCallExpr(t, `c.json({});`)
	if got := extractSecondNumberArg(call, src, "200"); got != "200" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractSecondNumberArg_CustomDefault(t *testing.T) {
	call, src := firstCallExpr(t, `c.body(null);`)
	if got := extractSecondNumberArg(call, src, "204"); got != "204" {
		t.Fatalf("got %q", got)
	}
}
