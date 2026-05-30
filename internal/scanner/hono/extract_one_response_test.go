//ff:func feature=scan type=test control=sequence topic=hono
//ff:what extractOneResponse 테스트
package hono

import "testing"

func firstCall(t *testing.T, src string) *fileInfo {
	t.Helper()
	return mustParse(t, []byte(src+"\n"))
}

func TestExtractOneResponse_Json(t *testing.T) {
	fi := firstCall(t, `c.json({ ok: true });`)
	call := findAllByType(fi.Root, "call_expression")[0]
	r := extractOneResponse(call, fi.Src)
	if r == nil || r.Kind != "json" || r.Status != "200" {
		t.Fatalf("got %+v", r)
	}
}

func TestExtractOneResponse_JsonStatus(t *testing.T) {
	fi := firstCall(t, `c.json({ ok: true }, 201);`)
	call := findAllByType(fi.Root, "call_expression")[0]
	r := extractOneResponse(call, fi.Src)
	if r == nil || r.Status != "201" {
		t.Fatalf("got %+v", r)
	}
}

func TestExtractOneResponse_Text(t *testing.T) {
	fi := firstCall(t, `c.text("hi");`)
	call := findAllByType(fi.Root, "call_expression")[0]
	r := extractOneResponse(call, fi.Src)
	if r == nil || r.Kind != "text" || r.Status != "200" {
		t.Fatalf("got %+v", r)
	}
}

func TestExtractOneResponse_Body(t *testing.T) {
	fi := firstCall(t, `c.body(null, 204);`)
	call := findAllByType(fi.Root, "call_expression")[0]
	r := extractOneResponse(call, fi.Src)
	if r == nil || r.Kind != "empty" || r.Status != "204" {
		t.Fatalf("got %+v", r)
	}
}

func TestExtractOneResponse_UnknownMethod(t *testing.T) {
	fi := firstCall(t, `c.render("x");`)
	call := findAllByType(fi.Root, "call_expression")[0]
	if r := extractOneResponse(call, fi.Src); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestExtractOneResponse_NotContext(t *testing.T) {
	// object identifier is not "c"
	fi := firstCall(t, `res.json({});`)
	call := findAllByType(fi.Root, "call_expression")[0]
	if r := extractOneResponse(call, fi.Src); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestExtractOneResponse_NoIdentifierObject(t *testing.T) {
	// `this.json` -> member_expression object is not an identifier -> obj nil
	fi := firstCall(t, `this.json({});`)
	call := findAllByType(fi.Root, "call_expression")[0]
	if r := extractOneResponse(call, fi.Src); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestExtractOneResponse_NoMemberExpr(t *testing.T) {
	// plain identifier call, no member_expression
	fi := firstCall(t, `foo();`)
	call := findAllByType(fi.Root, "call_expression")[0]
	if r := extractOneResponse(call, fi.Src); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
