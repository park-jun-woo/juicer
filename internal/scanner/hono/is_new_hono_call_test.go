//ff:func feature=scan type=test control=sequence topic=hono
//ff:what isNewHonoCall 테스트
package hono

import "testing"

func newExprOf(t *testing.T, src string) (*fileInfo, bool) {
	t.Helper()
	fi := mustParse(t, []byte(src+"\n"))
	news := findAllByType(fi.Root, "new_expression")
	if len(news) == 0 {
		return fi, false
	}
	return fi, isNewHonoCall(news[0], fi.Src)
}

func TestIsNewHonoCall_Hono(t *testing.T) {
	if _, ok := newExprOf(t, `const a = new Hono();`); !ok {
		t.Fatal("expected true for new Hono()")
	}
}

func TestIsNewHonoCall_OpenAPIHonoVariant(t *testing.T) {
	if _, ok := newExprOf(t, `const a = new OpenAPIHono();`); !ok {
		t.Fatal("expected true for new OpenAPIHono()")
	}
}

func TestIsNewHonoCall_Other(t *testing.T) {
	if _, ok := newExprOf(t, `const a = new Other();`); ok {
		t.Fatal("expected false for new Other()")
	}
}

func TestIsNewHonoCall_NotNewExpr(t *testing.T) {
	fi := mustParse(t, []byte(`foo();` + "\n"))
	call := findAllByType(fi.Root, "call_expression")[0]
	if isNewHonoCall(call, fi.Src) {
		t.Fatal("expected false for non-new_expression")
	}
}
