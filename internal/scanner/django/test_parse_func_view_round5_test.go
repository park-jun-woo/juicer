//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParseFuncView_Round5 테스트
package django

import "testing"

func TestParseFuncView_Round5(t *testing.T) {
	fi := newTestFileInfo(t, "@api_view(['GET'])\ndef ping(request):\n    return Response()\n")
	fn := djFirst(t, fi.root, "function_definition")
	fv := parseFuncView(fn, fi)
	if fv == nil {
		t.Fatal("expected func view")
	}
	if fv.name != "ping" {
		t.Errorf("name: %q", fv.name)
	}
}
