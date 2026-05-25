//ff:func feature=ratchet type=session control=sequence
//ff:what TestQueryExists_NoDir 테스트
package sqls

import (
	"testing"
)

func TestQueryExists_NoDir(t *testing.T) {
	if queryExists("/nonexistent/dir", "query") {
		t.Error("expected false for nonexistent dir")
	}
}
