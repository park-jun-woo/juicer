package sqls

import "testing"

func TestToQueryName_WithRepo(t *testing.T) {
	got := toQueryName("UserRepository.FindByID")
	if got != "UserFindByID" {
		t.Fatalf("expected UserFindByID, got %s", got)
	}
}

func TestToQueryName_NoDot(t *testing.T) {
	got := toQueryName("FindByID")
	if got != "FindByID" {
		t.Fatalf("expected FindByID, got %s", got)
	}
}
