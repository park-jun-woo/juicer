//ff:func feature=scan type=test control=sequence
//ff:what TestBuildEndpointIndex_Empty 테스트
package fiber

import "testing"

func TestBuildEndpointIndex_Empty(t *testing.T) {
	if len(buildEndpointIndex(nil)) != 0 {
		t.Error("expected empty map")
	}
}
