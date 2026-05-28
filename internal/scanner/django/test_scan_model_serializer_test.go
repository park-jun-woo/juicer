//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what ModelSerializer의 Meta.fields에서 필드명을 추출한다
package django

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractModelSerializerMetaFields(t *testing.T) {
	dir := t.TempDir()

	src := `from rest_framework import serializers

class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = ["id", "name", "email", "created_at"]
`
	os.WriteFile(filepath.Join(dir, "serializers.py"), []byte(src), 0o644)

	files := parseAllFiles(dir, []string{filepath.Join(dir, "serializers.py")})
	serializers := extractSerializers(files)

	si, ok := serializers["UserSerializer"]
	if !ok {
		t.Fatal("UserSerializer not found")
	}

	if len(si.fields) != 4 {
		for _, f := range si.fields {
			t.Logf("  field: %s (type=%s)", f.Name, f.Type)
		}
		t.Fatalf("expected 4 fields, got %d", len(si.fields))
	}

	expectedNames := []string{"id", "name", "email", "created_at"}
	for i, name := range expectedNames {
		if si.fields[i].Name != name {
			t.Errorf("field %d: expected %s, got %s", i, name, si.fields[i].Name)
		}
	}
}
