//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestSubstituteType 테스트
package quarkus

import "testing"

func TestSubstituteType(t *testing.T) {
	m := map[string]string{"T": "string"}
	if got := substituteType("T", m); got != "string" {
		t.Fatalf("plain: %q", got)
	}
	if got := substituteType("[]T", m); got != "[]string" {
		t.Fatalf("slice: %q", got)
	}
	if got := substituteType("array:T", m); got != "array:string" {
		t.Fatalf("array: %q", got)
	}
	if got := substituteType("List<T>", m); got != "List<string>" {
		t.Fatalf("generic: %q", got)
	}
	if got := substituteType("Other", m); got != "Other" {
		t.Fatalf("unmapped: %q", got)
	}
}
