//ff:func feature=scan type=test control=sequence
//ff:what ensureSchema 등록 로직 테스트
package scanner

import (
	"fmt"
	"testing"
)

func TestEnsureSchema(t *testing.T) {
	schemas := map[string]any{}

	// fields가 있으면 fieldsToSchema 결과를 등록한다
	ensureSchema("user", []Field{{Name: "id", JSON: "id", Type: "int"}}, schemas)
	if schemas["user"] == nil {
		t.Fatal("expected schema to be registered")
	}

	// 이미 등록된 schemaName은 덮어쓰지 않는다
	snapshot := fmt.Sprintf("%v", schemas["user"])
	ensureSchema("user", []Field{{Name: "name", JSON: "name", Type: "string"}}, schemas)
	if fmt.Sprintf("%v", schemas["user"]) != snapshot {
		t.Fatal("expected schema not to be overwritten")
	}

	// fields가 비어있으면 빈 object schema를 등록한다
	ensureSchema("empty", nil, schemas)
	m, ok := schemas["empty"].(map[string]any)
	if !ok || m["type"] != "object" {
		t.Fatalf("expected {type: object}, got %v", schemas["empty"])
	}
}
