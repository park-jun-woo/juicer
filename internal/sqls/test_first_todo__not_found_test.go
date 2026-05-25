//ff:func feature=sql type=parse control=sequence
//ff:what TestFirstTODO_NotFound 테스트
package sqls

import "testing"

func TestFirstTODO_NotFound(t *testing.T) {
	sess := &Session{Methods: []MethodStatus{{Status: "DONE"}}}
	if firstTODO(sess) != -1 {
		t.Fatal("expected -1")
	}
}
