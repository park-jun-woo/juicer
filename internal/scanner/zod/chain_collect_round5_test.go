//ff:func feature=scan type=test control=sequence topic=zod
//ff:what chain 수집/스키마 함수 테스트 (round5)
package zod

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

// findFirstCall returns the first call_expression node in the tree.
func findFirstCall(t *testing.T, root *sitter.Node) *sitter.Node {
	t.Helper()
	var found *sitter.Node
	walkNodes(root, func(n *sitter.Node) {
		if found == nil && n.Type() == "call_expression" {
			found = n
		}
	})
	if found == nil {
		t.Fatal("no call_expression found")
	}
	return found
}

func TestCollectChainMethods_Round5(t *testing.T) {
	root, src := parseTS(t, "const s = z.string().min(3).email();")
	call := findFirstCall(t, root)
	methods := CollectChainMethods(call, src)
	if len(methods) == 0 {
		t.Fatalf("expected chain methods, got %+v", methods)
	}
	// inner->outer order: first should be the base "string"
	names := make([]string, len(methods))
	for i, m := range methods {
		names[i] = m.Name
	}
	hasString, hasMin, hasEmail := false, false, false
	for _, n := range names {
		switch n {
		case "string":
			hasString = true
		case "min":
			hasMin = true
		case "email":
			hasEmail = true
		}
	}
	if !hasString || !hasMin || !hasEmail {
		t.Fatalf("missing chain methods: %v", names)
	}
}

func TestCollectChainMethodsRecursive_NilAndDefault_Round5(t *testing.T) {
	var methods []ChainMethod
	// nil node returns early
	collectChainMethodsRecursive(nil, nil, &methods)
	if len(methods) != 0 {
		t.Fatal("nil node should not append")
	}
	// a non call/member node (identifier) hits the default (no-op) branch
	root, src := parseTS(t, "const x = 1;")
	var idNode *sitter.Node
	walkNodes(root, func(n *sitter.Node) {
		if idNode == nil && n.Type() == "identifier" {
			idNode = n
		}
	})
	if idNode == nil {
		t.Fatal("no identifier")
	}
	collectChainMethodsRecursive(idNode, src, &methods)
	if len(methods) != 0 {
		t.Fatal("identifier should be no-op")
	}
}

func TestCollectChainFromMemberExpr_Round5(t *testing.T) {
	// z.string is a member_expression with object "z"
	root, src := parseTS(t, "const s = z.string;")
	var member *sitter.Node
	walkNodes(root, func(n *sitter.Node) {
		if member == nil && n.Type() == "member_expression" {
			member = n
		}
	})
	if member == nil {
		t.Fatal("no member_expression")
	}
	var methods []ChainMethod
	collectChainFromMemberExpr(member, src, &methods)
	if len(methods) != 1 || methods[0].Name != "string" {
		t.Fatalf("expected [string], got %+v", methods)
	}
}

func TestCollectChainFromCallExpr_Round5(t *testing.T) {
	root, src := parseTS(t, "const s = z.string().min(2);")
	call := findFirstCall(t, root)
	var methods []ChainMethod
	collectChainFromCallExpr(call, src, &methods)
	if len(methods) == 0 {
		t.Fatalf("expected methods, got %+v", methods)
	}
}

func TestResolveFunctionNode_Round5(t *testing.T) {
	root, _ := parseTS(t, "const s = z.string();")
	call := findFirstCall(t, root)
	fn := resolveFunctionNode(call)
	if fn == nil {
		t.Fatal("expected function node")
	}
	if fn.Type() != "member_expression" {
		t.Fatalf("expected member_expression, got %s", fn.Type())
	}
}

func TestBuildChainMethodFromProp_Round5(t *testing.T) {
	root, src := parseTS(t, "const s = z.string().min(3);")
	// find the .min(...) call: it's a call_expression whose function property is "min"
	var minCall *sitter.Node
	walkNodes(root, func(n *sitter.Node) {
		if n.Type() != "call_expression" {
			return
		}
		fn := resolveFunctionNode(n)
		if fn == nil || fn.Type() != "member_expression" {
			return
		}
		prop := fn.ChildByFieldName("property")
		if prop != nil && nodeText(prop, src) == "min" {
			minCall = n
		}
	})
	if minCall == nil {
		t.Fatal("no min call")
	}
	fn := resolveFunctionNode(minCall)
	prop := fn.ChildByFieldName("property")
	cm := buildChainMethodFromProp(minCall, prop, src)
	if cm.Name != "min" {
		t.Fatalf("name: got %q", cm.Name)
	}
	if len(cm.Args) != 1 || cm.Args[0] != "3" {
		t.Fatalf("args: got %+v", cm.Args)
	}
}

func TestCollectSchemaFromDecl_Round5(t *testing.T) {
	root, src := parseTS(t, "const UserSchema = z.object({});")
	var decl *sitter.Node
	walkNodes(root, func(n *sitter.Node) {
		if decl == nil && (n.Type() == "lexical_declaration" || n.Type() == "variable_declaration") {
			decl = n
		}
	})
	if decl == nil {
		t.Fatal("no declaration")
	}
	schemas := map[string]*sitter.Node{}
	CollectSchemaFromDecl(decl, src, schemas)
	if _, ok := schemas["UserSchema"]; !ok {
		t.Fatalf("UserSchema not collected: %+v", schemas)
	}
}
