//ff:func feature=scan type=test control=sequence topic=django
//ff:what Serializer 필드를 파싱하고 제약조건을 추출한다
package django

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractSerializerFields(t *testing.T) {
	dir := t.TempDir()

	src := `from rest_framework import serializers

class UserSerializer(serializers.Serializer):
    name = serializers.CharField(max_length=100)
    email = serializers.EmailField()
    age = serializers.IntegerField(min_value=0, max_value=150)
    is_active = serializers.BooleanField()
    role = serializers.ChoiceField(choices=["admin", "user"])
    avatar = serializers.FileField()
`
	os.WriteFile(filepath.Join(dir, "serializers.py"), []byte(src), 0o644)

	files := parseAllFiles(dir, []string{filepath.Join(dir, "serializers.py")})
	serializers := extractSerializers(files)

	si, ok := serializers["UserSerializer"]
	if !ok {
		t.Fatal("UserSerializer not found")
	}

	if len(si.fields) != 6 {
		for _, f := range si.fields {
			t.Logf("  field: %s (type=%s)", f.Name, f.Type)
		}
		t.Fatalf("expected 6 fields, got %d", len(si.fields))
	}

	// Check name field
	nameField := si.fields[0]
	if nameField.Name != "name" || nameField.Type != "string" {
		t.Errorf("field 0: expected name (string), got %s (%s)", nameField.Name, nameField.Type)
	}
	if nameField.MaxLength == nil || *nameField.MaxLength != 100 {
		t.Errorf("field 0: expected max_length=100")
	}

	// Check email field
	emailField := si.fields[1]
	if emailField.Name != "email" || emailField.Type != "string" {
		t.Errorf("field 1: expected email (string), got %s (%s)", emailField.Name, emailField.Type)
	}

	// Check age field
	ageField := si.fields[2]
	if ageField.Name != "age" || ageField.Type != "integer" {
		t.Errorf("field 2: expected age (integer), got %s (%s)", ageField.Name, ageField.Type)
	}
	if ageField.Minimum == nil || *ageField.Minimum != 0 {
		t.Errorf("field 2: expected min_value=0")
	}
	if ageField.Maximum == nil || *ageField.Maximum != 150 {
		t.Errorf("field 2: expected max_value=150")
	}

	// Check role field (ChoiceField)
	roleField := si.fields[4]
	if roleField.Name != "role" {
		t.Errorf("field 4: expected role, got %s", roleField.Name)
	}
	if len(roleField.Enum) != 2 {
		t.Errorf("field 4: expected 2 enum values, got %d", len(roleField.Enum))
	}
}
