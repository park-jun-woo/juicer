//ff:func feature=sql type=parse control=sequence
//ff:what TestExtract_ReadDirError 테스트
package sqls

import (
	"testing"
)

func TestExtract_ReadDirError(t *testing.T) {
	_, err := Extract("/nonexistent/dir")
	if err == nil {
		t.Error("expected error for non-existent dir")
	}
}
