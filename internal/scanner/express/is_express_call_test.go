//ff:func feature=scan type=test control=sequence topic=express
//ff:what isExpressCall: express() true / 비call / 다른 함수 false
package express

import "testing"

func TestIsExpressCall_True(t *testing.T) {
	fi := mustParse(t, []byte(`const a = express();`))
	if !isExpressCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected true")
	}
}

func TestIsExpressCall_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`const a = express;`))
	// pass the identifier node, not a call
	ids := findAllByType(fi.Root, "identifier")
	if len(ids) == 0 {
		t.Fatal("no identifier")
	}
	if isExpressCall(ids[0], fi.Src) {
		t.Fatal("expected false for non-call")
	}
}

func TestIsExpressCall_OtherFn(t *testing.T) {
	fi := mustParse(t, []byte(`const a = foo();`))
	if isExpressCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected false for foo()")
	}
}
