//ff:func feature=scan type=test control=sequence
//ff:what SchemaFromFields 필드 목록 → OpenAPI 스키마 객체 변환 테스트
package scanner

import "testing"

func TestSchemaFromFields(t *testing.T) {
	got := SchemaFromFields([]Field{{Name: "name", Type: "string"}})
	if got["type"] != "object" {
		t.Errorf("type: %v", got["type"])
	}
	props, ok := got["properties"].(map[string]any)
	if !ok || props["name"] == nil {
		t.Errorf("properties: %v", got["properties"])
	}
}
