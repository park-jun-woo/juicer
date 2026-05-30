//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestSplitGenericArgs 테스트
package spring

import "testing"

func TestSplitGenericArgs(t *testing.T) {
	got := splitGenericArgs("A, Map<B,C>, D")
	if len(got) != 3 || got[1] != "Map<B,C>" {
		t.Fatalf("got %v", got)
	}
}
