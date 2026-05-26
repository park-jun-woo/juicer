//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDtoFieldsToScannerFields_Empty 테스트
package nestjs

import "testing"

func TestDtoFieldsToScannerFields_Empty(t *testing.T) {
	result := dtoFieldsToScannerFields(nil)
	if len(result) != 0 {
		t.Fatalf("expected 0, got %d", len(result))
	}
}
