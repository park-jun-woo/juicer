//ff:func feature=scan type=test control=sequence
//ff:what NestJS Scan 스텁 테스트
package nestjs

import "testing"

func TestScan_Stub(t *testing.T) {
	_, err := Scan("/nonexistent")
	if err == nil {
		t.Fatal("expected error from stub")
	}
}
