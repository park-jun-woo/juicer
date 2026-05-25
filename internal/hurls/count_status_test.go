package hurls

import "testing"

func TestCountStatus(t *testing.T) {
	sess := &Session{Endpoints: []EndpointStatus{
		{ID: "GET /a", Status: "TODO"},
		{ID: "GET /b", Status: "DONE"},
		{ID: "GET /c", Status: "TODO"},
	}}
	if got := countStatus(sess, "TODO"); got != 2 {
		t.Fatalf("expected 2, got %d", got)
	}
	if got := countStatus(sess, "DONE"); got != 1 {
		t.Fatalf("expected 1, got %d", got)
	}
	if got := countStatus(sess, "SKIP"); got != 0 {
		t.Fatalf("expected 0, got %d", got)
	}
}
