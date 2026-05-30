//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractTypeArgContent_Single 테스트
package actix

import "testing"

func TestExtractTypeArgContent_Single(t *testing.T) {
	src := []byte(`fn f(x: web::Json<User>) {}`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	ta := firstTypeArguments(root)
	if ta == nil {
		t.Fatal("no type_arguments found")
	}
	if got := extractTypeArgContent(ta, src); got != "User" {
		t.Fatalf("extractTypeArgContent = %q, want User", got)
	}
}
