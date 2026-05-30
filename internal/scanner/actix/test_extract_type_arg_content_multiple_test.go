//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractTypeArgContent_Multiple 테스트
package actix

import "testing"

func TestExtractTypeArgContent_Multiple(t *testing.T) {

	src := []byte(`fn f(x: HashMap<String, i32>) {}`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	ta := firstTypeArguments(root)
	if ta == nil {
		t.Fatal("no type_arguments found")
	}
	got := extractTypeArgContent(ta, src)

	if got == "String" || got == "" {
		t.Fatalf("unexpected single-type result for multi-arg: %q", got)
	}
}
