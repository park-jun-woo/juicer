//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestExtractMethodBlockFromSwitch_Round5 테스트
package supafunc

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestExtractMethodBlockFromSwitch_Round5(t *testing.T) {
	src := []byte(`
serve(async (req) => {
  switch (req.method) {
    case "GET":
      return new Response("get", { status: 200 })
    case "POST":
      return new Response("post", { status: 201 })
    case "OPTIONS":
      return new Response(null, { status: 204 })
  }
})
`)
	fi := mustParse(t, src)
	// find the switch_statement node
	var sw *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if sw != nil {
			return
		}
		if n.Type() == "switch_statement" {
			sw = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(fi.Root)
	if sw == nil {
		t.Fatal("no switch statement")
	}
	result := map[string]*sitter.Node{}
	extractMethodBlockFromSwitch(sw, fi.Src, result)
	if result["GET"] == nil || result["POST"] == nil {
		t.Fatalf("expected GET and POST blocks")
	}
	if result["OPTIONS"] != nil {
		t.Fatal("OPTIONS should be skipped")
	}
}
