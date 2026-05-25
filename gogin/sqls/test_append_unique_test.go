//ff:func feature=sql type=parse control=sequence
//ff:what TestAppendUnique 테스트
package sqls

import (
	"testing"
)

func TestAppendUnique(t *testing.T) {
	slice := []string{"a", "b"}
	result := appendUnique(slice, "c")
	if len(result) != 3 {
		t.Errorf("expected 3, got %d", len(result))
	}

	result2 := appendUnique(result, "a")
	if len(result2) != 3 {
		t.Errorf("expected 3 (no duplicate), got %d", len(result2))
	}
}
