//ff:func feature=scan type=test control=sequence
//ff:what TestDeduplicateEndpoints_Empty 테스트
package scanner

import "testing"

func TestDeduplicateEndpoints_Empty(t *testing.T) {
	result := deduplicateEndpoints(nil)
	if len(result) != 0 {
		t.Fatal("expected empty")
	}
}

