//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestSplitGenericArgs 테스트
package quarkus

import "testing"

func TestSplitGenericArgs(t *testing.T) {
	got := splitGenericArgs("String, Map<String,Integer>, Long")
	if len(got) != 3 || got[1] != "Map<String,Integer>" {
		t.Fatalf("got %v", got)
	}
}
