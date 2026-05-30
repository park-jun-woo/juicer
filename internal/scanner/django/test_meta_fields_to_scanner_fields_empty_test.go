//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestMetaFieldsToScannerFields_Empty 테스트
package django

import "testing"

func TestMetaFieldsToScannerFields_Empty(t *testing.T) {
	if f := metaFieldsToScannerFields(nil); f != nil {
		t.Fatalf("expected nil, got %+v", f)
	}
}
