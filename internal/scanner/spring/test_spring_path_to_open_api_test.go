//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestSpringPathToOpenAPI 테스트
package spring

import "testing"

func TestSpringPathToOpenAPI(t *testing.T) {
	if springPathToOpenAPI("/users/{id}") != "/users/{id}" {
		t.Fatal("path")
	}
}
