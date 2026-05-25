//ff:func feature=hurl type=parse control=sequence
//ff:what TestBuildEndpointStatuses_Empty 테스트
package hurls

import "testing"

func TestBuildEndpointStatuses_Empty(t *testing.T) {
	statuses := buildEndpointStatuses(nil)
	if len(statuses) != 0 {
		t.Fatalf("expected 0, got %d", len(statuses))
	}
}
