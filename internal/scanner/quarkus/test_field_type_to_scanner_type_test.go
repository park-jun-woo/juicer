//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestFieldTypeToScannerType 테스트
package quarkus

import "testing"

func TestFieldTypeToScannerType(t *testing.T) {
	if got := fieldTypeToScannerType("String"); got != "string" {
		t.Fatalf("string: %q", got)
	}
	if got := fieldTypeToScannerType("UUID"); got != "string:uuid" {
		t.Fatalf("uuid: %q", got)
	}
	if got := fieldTypeToScannerType("List<String>"); got != "array:string" {
		t.Fatalf("array: %q", got)
	}
}
