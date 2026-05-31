//ff:func feature=scan type=test control=sequence
//ff:what TestFieldToProperty_ArrayObjectRef — 배열 요소가 객체 타입이면 $ref를 생성한다
package scanner

import "testing"

func TestFieldToProperty_ArrayObjectRef(t *testing.T) {
	f := Field{Name: "Photos", Type: "[]Photo"}
	prop := fieldToProperty(f)
	if prop["type"] != "array" {
		t.Fatalf("expected type=array, got %v", prop["type"])
	}
	items, ok := prop["items"].(map[string]any)
	if !ok {
		t.Fatalf("expected items map, got %T", prop["items"])
	}
	ref, ok := items["$ref"]
	if !ok {
		t.Fatal("expected $ref in items")
	}
	if ref != "#/components/schemas/Photo" {
		t.Errorf("expected $ref=#/components/schemas/Photo, got %v", ref)
	}
}
