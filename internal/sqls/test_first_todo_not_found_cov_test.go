//ff:func feature=sql type=test control=sequence
//ff:what TestFirstTODO_NotFoundCov 테스트
package sqls

import "testing"

func TestFirstTODO_NotFoundCov(t *testing.T) {
	sess := &Session{Methods: []MethodStatus{
		{Status: "DONE"}, {Status: "DONE"},
	}}
	if firstTODO(sess) != -1 {
		t.Fatal("expected -1")
	}
}
