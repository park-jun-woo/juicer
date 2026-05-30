//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCheckParamTypeAnnotation_ExpressRouter 테스트
package express

import "testing"

func TestCheckParamTypeAnnotation_ExpressRouter(t *testing.T) {
	fi := mustParse(t, []byte("function f(r: express.Router) {}\n"))
	param := firstParamOfType(fi.Root, "required_parameter")
	routers := map[string]bool{}
	checkParamTypeAnnotation(param, fi.Src, routers)
	if !routers["r"] {
		t.Fatalf("expected r registered, got %v", routers)
	}
}
