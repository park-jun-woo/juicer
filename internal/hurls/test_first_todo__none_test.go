//ff:func feature=hurl type=parse control=sequence
//ff:what TestFirstTODO_None 테스트
package hurls

import "testing"

func TestFirstTODO_None(t *testing.T) {
	sess := &Session{Endpoints: []EndpointStatus{
		{ID: "GET /a", Status: "DONE"},
	}}
	if got := firstTODO(sess); got != -1 {
		t.Fatalf("expected -1, got %d", got)
	}
}
