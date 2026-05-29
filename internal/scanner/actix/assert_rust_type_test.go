//ff:func feature=scan type=test control=sequence topic=actix
//ff:what assertRustType — rustTypeToOpenAPI 결과를 기대값과 비교한다
package actix

import "testing"

func assertRustType(t *testing.T, input, wantType, wantFormat, wantItems string) {
	t.Helper()
	got := rustTypeToOpenAPI(input)
	if got.Type != wantType {
		t.Errorf("type: want %s, got %s", wantType, got.Type)
	}
	if got.Format != wantFormat {
		t.Errorf("format: want %s, got %s", wantFormat, got.Format)
	}
	if got.Items != wantItems {
		t.Errorf("items: want %s, got %s", wantItems, got.Items)
	}
}
