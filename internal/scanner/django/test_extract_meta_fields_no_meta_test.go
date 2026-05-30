//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractMetaFields_NoMeta 테스트
package django

import "testing"

func TestExtractMetaFields_NoMeta(t *testing.T) {
	src := `
class UserSerializer(ModelSerializer):
    other = 1
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	body := firstBlock(root)
	if f := extractMetaFields(body, []byte(src)); f != nil {
		t.Fatalf("expected nil without Meta, got %v", f)
	}
}
