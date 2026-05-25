package sqls

import "testing"

func TestCountStatus_Basic(t *testing.T) {
	sess := &Session{Methods: []MethodStatus{
		{Status: "TODO"}, {Status: "DONE"}, {Status: "TODO"},
	}}
	if countStatus(sess, "TODO") != 2 {
		t.Fatal("expected 2")
	}
}
