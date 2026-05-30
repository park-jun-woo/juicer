//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectParamRoutersFromType_FuncAndArrow 테스트
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
