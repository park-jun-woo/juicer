//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what TestExtractSerializerFields_Multi 테스트
package django

import "testing"

func TestExtractSerializerFields_Multi(t *testing.T) {
	src := `
class UserSerializer(Serializer):
    name = serializers.CharField(max_length=100)
    age = serializers.IntegerField()
    helper = 42
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	body := firstBlock(root)
	if body == nil {
		t.Fatal("no class body")
	}
	fields := extractSerializerFields(body, []byte(src))
	if len(fields) < 2 {
		t.Fatalf("expected at least 2 fields, got %d: %+v", len(fields), fields)
	}
	names := map[string]bool{}
	for _, f := range fields {
		names[f.Name] = true
	}
	if !names["name"] || !names["age"] {
		t.Errorf("expected name and age fields, got %+v", fields)
	}
}
