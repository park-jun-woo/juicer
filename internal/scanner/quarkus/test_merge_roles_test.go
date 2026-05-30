//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestMergeRoles 테스트
package quarkus

import "testing"

func TestMergeRoles(t *testing.T) {
	if got := mergeRoles([]string{"a"}, []string{"b"}); got[0] != "b" {
		t.Fatal("method wins")
	}
	if got := mergeRoles([]string{"a"}, nil); got[0] != "a" {
		t.Fatal("class fallback")
	}
}
