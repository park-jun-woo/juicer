//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractMetaFields 테스트
package django

import "testing"

func TestExtractMetaFields(t *testing.T) {
	src := `
class UserSerializer(ModelSerializer):
    class NotMeta:
        fields = ['ignored']
    class Meta:
        model = User
        fields = ['id', 'name']
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	body := firstBlock(root)
	if body == nil {
		t.Fatal("no class body")
	}
	fields := extractMetaFields(body, []byte(src))
	if len(fields) != 2 || fields[0] != "id" || fields[1] != "name" {
		t.Fatalf("expected [id name] from Meta, got %v", fields)
	}
}
