//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractApiViewDecorator_Round5 테스트
package django

import "testing"

func TestExtractApiViewDecorator_Round5(t *testing.T) {
	src := []byte("@api_view(['GET', 'POST'])\ndef v(request):\n    return Response()\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	fn := djFirst(t, root, "function_definition")
	methods := extractAPIViewDecorator(fn, src)
	if len(methods) != 2 {
		t.Fatalf("expected 2 methods, got %v", methods)
	}
}
