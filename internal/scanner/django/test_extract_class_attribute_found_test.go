//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractClassAttribute_Found 테스트
package django

import "testing"

func TestExtractClassAttribute_Found(t *testing.T) {
	src := `
class UserViewSet(ModelViewSet):
    queryset = User.objects.all()
    serializer_class = UserSerializer
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	body := firstBlock(root)
	if body == nil {
		t.Fatal("no class body")
	}
	got := extractClassAttribute(body, "serializer_class", []byte(src))
	if got != "UserSerializer" {
		t.Fatalf("got %q, want UserSerializer", got)
	}
}
