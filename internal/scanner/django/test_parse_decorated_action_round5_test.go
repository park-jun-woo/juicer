//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParseDecoratedAction_Round5 테스트
package django

import "testing"

func TestParseDecoratedAction_Round5(t *testing.T) {
	src := []byte("@action(detail=True, methods=['post'])\ndef activate(self, request, pk=None):\n    return Response()\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	dec := djFirst(t, root, "decorated_definition")
	action := parseDecoratedAction(dec, src)
	if action == nil {
		t.Fatal("expected action")
	}
}
