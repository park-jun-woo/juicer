//ff:func feature=scan type=test control=sequence topic=express
//ff:what isResStatusCall: res.status() true / 비call / 비멤버 / prop불일치 / 비res객체
package express

import "testing"

func TestIsResStatusCall_True(t *testing.T) {
	fi := mustParse(t, []byte(`res.status(200);`))
	if !isResStatusCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected true")
	}
}

func TestIsResStatusCall_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`res;`))
	ids := findAllByType(fi.Root, "identifier")
	if isResStatusCall(ids[0], fi.Src) {
		t.Fatal("expected false")
	}
}

func TestIsResStatusCall_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`status(200);`))
	if isResStatusCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected false")
	}
}

func TestIsResStatusCall_PropMismatch(t *testing.T) {
	fi := mustParse(t, []byte(`res.json({});`))
	if isResStatusCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected false")
	}
}

func TestIsResStatusCall_NonResObject(t *testing.T) {
	fi := mustParse(t, []byte(`other.status(200);`))
	if isResStatusCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected false")
	}
}
