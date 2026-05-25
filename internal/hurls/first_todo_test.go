package hurls

import "testing"

func TestFirstTODO_Found(t *testing.T) {
	sess := &Session{Endpoints: []EndpointStatus{
		{ID: "GET /a", Status: "DONE"},
		{ID: "GET /b", Status: "TODO"},
	}}
	if got := firstTODO(sess); got != 1 {
		t.Fatalf("expected 1, got %d", got)
	}
}

func TestFirstTODO_None(t *testing.T) {
	sess := &Session{Endpoints: []EndpointStatus{
		{ID: "GET /a", Status: "DONE"},
	}}
	if got := firstTODO(sess); got != -1 {
		t.Fatalf("expected -1, got %d", got)
	}
}

func TestFirstTODO_Empty(t *testing.T) {
	sess := &Session{}
	if got := firstTODO(sess); got != -1 {
		t.Fatalf("expected -1, got %d", got)
	}
}
