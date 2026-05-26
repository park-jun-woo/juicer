//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what tryList 테스트
package fastapi

import "testing"

func TestTryList(t *testing.T) {
	oa, ok := tryList("list[str]")
	if !ok || oa.Type != "array" || oa.Items != "str" {
		t.Fatalf("got %v, %v", oa, ok)
	}
	oa2, ok := tryList("List[int]")
	if !ok || oa2.Type != "array" || oa2.Items != "int" {
		t.Fatalf("got %v, %v", oa2, ok)
	}
	_, ok = tryList("str")
	if ok {
		t.Fatal("should not match plain type")
	}
}
