package sqls

import "testing"

func TestFirstTODO_Found(t *testing.T) {
	sess := &Session{Methods: []MethodStatus{
		{Status: "DONE"}, {Status: "TODO"},
	}}
	if firstTODO(sess) != 1 {
		t.Fatal("expected 1")
	}
}

func TestFirstTODO_NotFound(t *testing.T) {
	sess := &Session{Methods: []MethodStatus{{Status: "DONE"}}}
	if firstTODO(sess) != -1 {
		t.Fatal("expected -1")
	}
}
