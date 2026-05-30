//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestSubstituteType 테스트
package spring

import "testing"

func TestSubstituteType(t *testing.T) {
	m := map[string]string{"T": "string"}
	if substituteType("T", m) != "string" || substituteType("[]T", m) != "[]string" {
		t.Fatal("substitute")
	}
	if substituteType("List<T>", m) != "List<string>" {
		t.Fatal("generic substitute")
	}
}
