//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCheckParamTypeAnnotation_OtherType 테스트
package express

import "testing"

func TestCheckParamTypeAnnotation_OtherType(t *testing.T) {
	fi := mustParse(t, []byte("function f(x: number) {}\n"))
	param := firstParamOfType(fi.Root, "required_parameter")
	routers := map[string]bool{}
	checkParamTypeAnnotation(param, fi.Src, routers)
	if len(routers) != 0 {
		t.Fatalf("expected no routers for non-Router type, got %v", routers)
	}
}
