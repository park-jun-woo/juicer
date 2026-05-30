//ff:func feature=scan type=test control=sequence topic=express
//ff:what nodeText: 노드의 소스 텍스트 반환
package express

import "testing"

func TestNodeText(t *testing.T) {
	fi := mustParse(t, []byte(`const userVar = 1;`))
	ids := findAllByType(fi.Root, "identifier")
	if len(ids) == 0 {
		t.Fatal("no identifier")
	}
	if got := nodeText(ids[0], fi.Src); got != "userVar" {
		t.Fatalf("got %q", got)
	}
}
