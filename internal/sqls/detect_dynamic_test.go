package sqls

import "testing"

func TestDetectDynamic_Nil(t *testing.T) {
	if detectDynamic(nil) {
		t.Fatal("expected false")
	}
}
