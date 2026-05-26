//ff:func feature=scan type=test control=sequence
//ff:what TestBuildRegisteredSet_Empty 테스트
package scanner

import "testing"

func TestBuildRegisteredSet_Empty(t *testing.T) {
	sr := &ScanResult{}
	set := buildRegisteredSet(sr)
	if len(set) != 0 {
		t.Fatal("expected empty set")
	}
}
