//ff:func feature=scan type=test control=sequence topic=django
//ff:what extractClassAttribute — 클래스 body 단순 속성 추출 분기를 검증
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

func TestExtractClassAttribute_NotFound(t *testing.T) {
	src := `
class C:
    other = X
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	body := firstBlock(root)
	if got := extractClassAttribute(body, "serializer_class", []byte(src)); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
