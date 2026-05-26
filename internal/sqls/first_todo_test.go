//ff:func feature=sql type=test control=sequence
//ff:what TestFirstTODO_Found 테스트
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
