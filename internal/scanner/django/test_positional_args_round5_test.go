//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestPositionalArgs_Round5 테스트
package django

import "testing"

func TestPositionalArgs_Round5(t *testing.T) {
	src := []byte("path('a/', AView.as_view(), name='a')\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := djFirst(t, root, "argument_list")
	pos := positionalArgs(args)

	if len(pos) != 2 {
		t.Fatalf("expected 2 positional args, got %d", len(pos))
	}
}
