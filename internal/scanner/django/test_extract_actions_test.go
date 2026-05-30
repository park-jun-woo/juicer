//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractActions 테스트
package django

import "testing"

func TestExtractActions(t *testing.T) {
	src := `
class UserViewSet(ModelViewSet):
    @action(detail=True, methods=['post'])
    def set_password(self, request, pk=None):
        return Response()

    def plain(self):
        pass
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	body := firstBlock(root)
	if body == nil {
		t.Fatal("no class body block")
	}
	actions := extractActions(body, []byte(src), "views.py")
	if len(actions) != 1 {
		t.Fatalf("expected 1 @action method, got %d", len(actions))
	}
	if actions[0].name != "set_password" {
		t.Errorf("name = %q, want set_password", actions[0].name)
	}
}
