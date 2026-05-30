//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what extractBodyMemberAccess / extractMethodBlockFromSwitch 테스트 (round5)
package supafunc

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestExtractBodyMemberAccess_Round5(t *testing.T) {
	src := []byte(`
serve(async (req) => {
  const body = await req.json()
  const name = body.name
  const age = body.age
  return new Response("ok", { status: 200 })
})
`)
	fi := mustParse(t, src)
	body, _ := findServeCallback(fi)
	if body == nil {
		t.Fatal("no callback body")
	}
	fields := extractBodyMemberAccess(body, fi.Src)
	if len(fields) == 0 {
		t.Fatalf("expected member-access fields, got %v", fields)
	}
}

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
