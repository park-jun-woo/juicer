//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestResolveViewSetMethods_Round5 테스트
package django

import "testing"

func TestResolveViewSetMethods_Round5(t *testing.T) {
	methods := resolveViewSetMethods([]string{"ModelViewSet"})
	if len(methods) == 0 {
		t.Fatal("ModelViewSet should expand to CRUD methods")
	}
}
