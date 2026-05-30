//ff:func feature=scan type=test control=sequence topic=express
//ff:what isResMethodCall: res.json / res.x불일치 / status체인 / 비call / 비멤버 분기
package express

import "testing"

func TestIsResMethodCall_DirectRes(t *testing.T) {
	fi := mustParse(t, []byte(`res.json({});`))
	m, ok := isResMethodCall(firstCallExpr(t, fi), fi.Src)
	if !ok || m != "json" {
		t.Fatalf("got %q,%v", m, ok)
	}
}

func TestIsResMethodCall_ResNonMethod(t *testing.T) {
	fi := mustParse(t, []byte(`res.render('v');`))
	if _, ok := isResMethodCall(firstCallExpr(t, fi), fi.Src); ok {
		t.Fatal("expected false for render")
	}
}

func TestIsResMethodCall_StatusChain(t *testing.T) {
	fi := mustParse(t, []byte(`res.status(201).json({});`))
	m, ok := isResMethodCall(outermostCall(fi), fi.Src)
	if !ok || m != "json" {
		t.Fatalf("got %q,%v", m, ok)
	}
}

func TestIsResMethodCall_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`res;`))
	ids := findAllByType(fi.Root, "identifier")
	if _, ok := isResMethodCall(ids[0], fi.Src); ok {
		t.Fatal("expected false for non-call")
	}
}

func TestIsResMethodCall_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`json({});`))
	if _, ok := isResMethodCall(firstCallExpr(t, fi), fi.Src); ok {
		t.Fatal("expected false for plain call")
	}
}

func TestIsResMethodCall_NonResObject(t *testing.T) {
	fi := mustParse(t, []byte(`other.json({});`))
	if _, ok := isResMethodCall(firstCallExpr(t, fi), fi.Src); ok {
		t.Fatal("expected false for non-res object")
	}
}

func TestIsResMethodCall_CallObjNotStatus(t *testing.T) {
	// object is a call but not res.status() -> falls through to final false
	fi := mustParse(t, []byte(`foo().json({});`))
	if _, ok := isResMethodCall(outermostCall(fi), fi.Src); ok {
		t.Fatal("expected false")
	}
}
