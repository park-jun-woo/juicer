//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestNodeText 테스트
package hono

import "testing"

func TestNodeText(t *testing.T) {
	fi := mustParse(t, []byte(`hello;`+"\n"))
	id := findAllByType(fi.Root, "identifier")[0]
	if got := nodeText(id, fi.Src); got != "hello" {
		t.Fatalf("got %q", got)
	}
}
