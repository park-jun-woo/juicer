//ff:func feature=scan type=test control=sequence topic=express
//ff:what isExpressRouterCall: express.Router() true / 비call / 비멤버 / 비obj / 불일치
package express

import "testing"

func TestIsExpressRouterCall_True(t *testing.T) {
	fi := mustParse(t, []byte(`const r = express.Router();`))
	if !isExpressRouterCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected true")
	}
}

func TestIsExpressRouterCall_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`const r = express;`))
	ids := findAllByType(fi.Root, "identifier")
	if isExpressRouterCall(ids[0], fi.Src) {
		t.Fatal("expected false for non-call")
	}
}

func TestIsExpressRouterCall_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`const r = express();`))
	if isExpressRouterCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected false for plain call")
	}
}

func TestIsExpressRouterCall_NoObjIdentifier(t *testing.T) {
	// member object is a member_expression -> no direct identifier
	fi := mustParse(t, []byte(`const r = a.b.Router();`))
	if isExpressRouterCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected false")
	}
}

func TestIsExpressRouterCall_Mismatch(t *testing.T) {
	fi := mustParse(t, []byte(`const r = foo.Bar();`))
	if isExpressRouterCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected false")
	}
}
