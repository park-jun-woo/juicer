//ff:func feature=scan type=test control=sequence topic=django
//ff:what metaFieldsToScannerFields — Meta.fields → scanner.Field 변환을 검증
package django

import "testing"

func TestMetaFieldsToScannerFields(t *testing.T) {
	fields := metaFieldsToScannerFields([]string{"id", "name"})
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(fields))
	}
	if fields[0].Name != "id" || fields[0].Type != "string" || fields[0].JSON != "id" {
		t.Errorf("unexpected field[0]: %+v", fields[0])
	}
}

func TestMetaFieldsToScannerFields_Empty(t *testing.T) {
	if f := metaFieldsToScannerFields(nil); f != nil {
		t.Fatalf("expected nil, got %+v", f)
	}
}
