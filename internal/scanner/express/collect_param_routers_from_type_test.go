//ff:func feature=scan type=test control=sequence topic=express
//ff:what collectParamRoutersFromType: arrow/function 선언 양쪽 분기 검증
package express

import "testing"

func TestCollectParamRoutersFromType_FuncAndArrow(t *testing.T) {
	src := []byte(`function fd(a: Router) {}
const af = (b: express.Router) => {};
`)
	fi := mustParse(t, src)
	routers := map[string]bool{}
	collectParamRoutersFromType(fi, routers)
	if !routers["a"] {
		t.Errorf("expected 'a' from function_declaration, got %v", routers)
	}
	if !routers["b"] {
		t.Errorf("expected 'b' from arrow_function, got %v", routers)
	}
}

func TestCollectParamRoutersFromType_None(t *testing.T) {
	src := []byte(`function fd(a: number) {}`)
	fi := mustParse(t, src)
	routers := map[string]bool{}
	collectParamRoutersFromType(fi, routers)
	if len(routers) != 0 {
		t.Errorf("expected none, got %v", routers)
	}
}
