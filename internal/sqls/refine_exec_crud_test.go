package sqls

import "testing"

func TestRefineExecCRUD_Insert(t *testing.T) {
	got := refineExecCRUD([]string{"INSERT INTO users"}, nil)
	if got != "INSERT" {
		t.Fatalf("expected INSERT, got %s", got)
	}
}

func TestRefineExecCRUD_NoSQL(t *testing.T) {
	got := refineExecCRUD(nil, nil)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
