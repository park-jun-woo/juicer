//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractGenericInner 테스트
package spring

import "testing"

func TestExtractGenericInner(t *testing.T) {
	if extractGenericInner("List<UserDto>") != "UserDto" || extractGenericInner("String") != "" {
		t.Fatal("generic inner")
	}
}
