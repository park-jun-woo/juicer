//ff:func feature=scan type=test topic=flask control=sequence
//ff:what attributeBaseKind request.form/request.json/기타 분류 테스트
package flask

import "testing"

func TestAttributeBaseKind(t *testing.T) {
	check := func(code, want string) {
		src := []byte(code)
		root, err := parsePython(src)
		if err != nil {
			t.Fatal(err)
		}
		attrs := findAllByType(root, "attribute")
		got := ""
		for _, a := range attrs {
			if k := attributeBaseKind(a, src); k != "" {
				got = k
			}
		}
		if got != want {
			t.Errorf("%q -> %q, want %q", code, got, want)
		}
	}
	check("x = request.form['a']\n", "form")
	check("x = request.json['a']\n", "json")
	check("x = request.args['a']\n", "")
}
