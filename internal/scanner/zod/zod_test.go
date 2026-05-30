//ff:func feature=scan type=test control=sequence topic=zod
//ff:what zod 패키지 전반 테스트 (Apply/Parse/Collect/Find/Resolve)
package zod

import (
	"context"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
	typescript "github.com/smacker/go-tree-sitter/typescript/typescript"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func parseTS(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	parser := sitter.NewParser()
	parser.SetLanguage(typescript.GetLanguage())
	tree, err := parser.ParseCtx(context.Background(), nil, b)
	if err != nil {
		t.Fatal(err)
	}
	return tree.RootNode(), b
}

func TestAppendValidate(t *testing.T) {
	if appendValidate("", "email") != "email" {
		t.Fatal("empty")
	}
	if appendValidate("required", "email") != "required,email" {
		t.Fatal("append")
	}
}

func TestUnquoteTS(t *testing.T) {
	if unquoteTS(`"x"`) != "x" || unquoteTS("'y'") != "y" || unquoteTS("`z`") != "z" || unquoteTS("a") != "a" {
		t.Fatal("unquote")
	}
}

func TestParseIntArg(t *testing.T) {
	if p := parseIntArg("5"); p == nil || *p != 5 {
		t.Fatal("5")
	}
	if parseIntArg("x") != nil {
		t.Fatal("invalid")
	}
}

func TestApplyMin(t *testing.T) {
	f := &scanner.Field{Type: "string"}
	ApplyMin(f, ChainMethod{Name: "min", Args: []string{"3"}})
	if f.MinLength == nil || *f.MinLength != 3 {
		t.Fatalf("string min: %v", f.MinLength)
	}
	f2 := &scanner.Field{Type: "number"}
	ApplyMin(f2, ChainMethod{Name: "min", Args: []string{"1"}})
	if f2.Minimum == nil || *f2.Minimum != 1 {
		t.Fatalf("number min: %v", f2.Minimum)
	}
	f3 := &scanner.Field{Type: "string"}
	ApplyMin(f3, ChainMethod{Name: "min"})
	if f3.MinLength != nil {
		t.Fatal("no args should be ignored")
	}
}

func TestApplyMax(t *testing.T) {
	f := &scanner.Field{Type: "string"}
	ApplyMax(f, ChainMethod{Name: "max", Args: []string{"10"}})
	if f.MaxLength == nil || *f.MaxLength != 10 {
		t.Fatalf("got %v", f.MaxLength)
	}
	f2 := &scanner.Field{Type: "number"}
	ApplyMax(f2, ChainMethod{Name: "max", Args: []string{"99"}})
	if f2.Maximum == nil || *f2.Maximum != 99 {
		t.Fatalf("got %v", f2.Maximum)
	}
}

func TestApplyMethod(t *testing.T) {
	cases := []struct {
		name  string
		args  []string
		check func(f scanner.Field) bool
	}{
		{"string", nil, func(f scanner.Field) bool { return f.Type == "string" }},
		{"number", nil, func(f scanner.Field) bool { return f.Type == "number" }},
		{"int", nil, func(f scanner.Field) bool { return f.Type == "integer" }},
		{"boolean", nil, func(f scanner.Field) bool { return f.Type == "boolean" }},
		{"email", nil, func(f scanner.Field) bool { return f.Validate == "email" }},
		{"url", nil, func(f scanner.Field) bool { return f.Validate == "uri" }},
		{"uuid", nil, func(f scanner.Field) bool { return f.Validate == "uuid" }},
		{"optional", nil, func(f scanner.Field) bool { return f.Nullable }},
		{"nullable", nil, func(f scanner.Field) bool { return f.Nullable }},
		{"array", nil, func(f scanner.Field) bool { return f.Type == "array" }},
		{"object", nil, func(f scanner.Field) bool { return f.Type == "object" }},
		{"enum", []string{"a", "b"}, func(f scanner.Field) bool { return f.Type == "string" && len(f.Enum) == 2 }},
	}
	for _, c := range cases {
		f := scanner.Field{}
		ApplyMethod(&f, ChainMethod{Name: c.name, Args: c.args})
		if !c.check(f) {
			t.Errorf("ApplyMethod(%q) failed: %+v", c.name, f)
		}
	}
}

func TestFindPathParamIndex(t *testing.T) {
	params := []scanner.Param{{Name: "id"}, {Name: "slug"}}
	if findPathParamIndex(params, "slug") != 1 {
		t.Fatal("found")
	}
	if findPathParamIndex(params, "missing") != -1 {
		t.Fatal("not found")
	}
}

func TestApplyParamValidator(t *testing.T) {
	req := &scanner.Request{PathParams: []scanner.Param{{Name: "id", Type: "string"}}}
	// existing param type updated
	ApplyParamValidator(req, []scanner.Field{{Name: "id", Type: "integer"}})
	if req.PathParams[0].Type != "integer" {
		t.Fatalf("update: %+v", req.PathParams)
	}
	// new param appended
	changed := ApplyParamValidator(req, []scanner.Field{{Name: "extra", Type: "string"}})
	if !changed || len(req.PathParams) != 2 {
		t.Fatalf("append: %+v", req.PathParams)
	}
}

func TestFindAllByTypeAndChild(t *testing.T) {
	root, _ := parseTS(t, `a(); b();`)
	if len(findAllByType(root, "call_expression")) != 2 {
		t.Fatal("findAllByType")
	}
	calls := findAllByType(root, "call_expression")
	if findChildByType(calls[0], "arguments") == nil {
		t.Fatal("findChildByType")
	}
	if findChildByType(calls[0], "object") != nil {
		t.Fatal("nil expected")
	}
}

func TestChildrenOfType(t *testing.T) {
	root, _ := parseTS(t, `const o = { a: 1, b: 2 };`)
	objs := findAllByType(root, "object")
	if len(childrenOfType(objs[0], "pair")) != 2 {
		t.Fatal("children")
	}
}

func TestCollectArgNodes(t *testing.T) {
	root, _ := parseTS(t, `f("a", b, 1);`)
	args := findAllByType(root, "arguments")[0]
	nodes := collectArgNodes(args)
	if len(nodes) != 3 {
		t.Fatalf("got %d", len(nodes))
	}
}

func TestExtractArrayStringValues(t *testing.T) {
	root, src := parseTS(t, `const x = ['a', 'b', 'c'];`)
	arr := findAllByType(root, "array")[0]
	vals := extractArrayStringValues(arr, src)
	if len(vals) != 3 || vals[0] != "a" {
		t.Fatalf("got %v", vals)
	}
}

func TestCollectStringArgs(t *testing.T) {
	root, src := parseTS(t, `z.enum(['a', 'b']);`)
	args := findAllByType(root, "arguments")[0]
	got := collectStringArgs(args, src)
	if len(got) != 2 {
		t.Fatalf("got %v", got)
	}
}

func TestIsObjectCall(t *testing.T) {
	root, src := parseTS(t, `z.object({ a: z.string() });`)
	calls := findAllByType(root, "call_expression")
	found := false
	for _, c := range calls {
		if IsObjectCall(c, src) {
			found = true
		}
	}
	if !found {
		t.Fatal("expected z.object call")
	}
	root2, src2 := parseTS(t, `other.foo();`)
	calls2 := findAllByType(root2, "call_expression")
	if IsObjectCall(calls2[0], src2) {
		t.Fatal("non-z call")
	}
}

func TestContainsCall(t *testing.T) {
	root, src := parseTS(t, `const s = z.string().min(1);`)
	if !ContainsCall(root, src) {
		t.Fatal("expected z call")
	}
	root2, src2 := parseTS(t, `const x = 1;`)
	if ContainsCall(root2, src2) {
		t.Fatal("no z call")
	}
}

func TestParseChainAndProperties(t *testing.T) {
	root, src := parseTS(t, `const s = z.object({ name: z.string().min(1), age: z.number().int() });`)
	calls := findAllByType(root, "call_expression")
	var objNode *sitter.Node
	for _, c := range calls {
		if IsObjectCall(c, src) {
			args := findChildByType(c, "arguments")
			objNode = findChildByType(args, "object")
			break
		}
	}
	if objNode == nil {
		t.Fatal("no object node")
	}
	fields := ParseObjectProperties(objNode, src)
	if len(fields) != 2 {
		t.Fatalf("got %d fields", len(fields))
	}
	if fields[0].Name != "name" || fields[0].Type != "string" {
		t.Fatalf("name field: %+v", fields[0])
	}
	if fields[0].MinLength == nil || *fields[0].MinLength != 1 {
		t.Fatalf("min: %v", fields[0].MinLength)
	}
}

func TestCollectSchemas(t *testing.T) {
	root, src := parseTS(t, `
const userSchema = z.object({ name: z.string() });
const x = 5;
`)
	schemas := CollectSchemas(root, src)
	if _, ok := schemas["userSchema"]; !ok {
		t.Fatalf("expected userSchema, got %v keys", len(schemas))
	}
}

func TestFindObjectCalls(t *testing.T) {
	root, src := parseTS(t, `z.object({ a: z.string() });`)
	calls := FindObjectCalls(root, src)
	if len(calls) == 0 {
		t.Fatal("expected object calls")
	}
}

func TestResolveValidatorFields_SchemaName(t *testing.T) {
	root, src := parseTS(t, `const userSchema = z.object({ name: z.string() });`)
	schemas := CollectSchemas(root, src)
	v := ValidatorInfo{Target: "json", SchemaName: "userSchema"}
	fields := ResolveValidatorFields(v, schemas, src, map[string][]byte{"userSchema": src})
	if len(fields) != 1 || fields[0].Name != "name" {
		t.Fatalf("got %+v", fields)
	}
}

func TestApplyValidator_JSON(t *testing.T) {
	root, src := parseTS(t, `const userSchema = z.object({ name: z.string() });`)
	schemas := CollectSchemas(root, src)
	req := &scanner.Request{}
	v := ValidatorInfo{Target: "json", SchemaName: "userSchema"}
	ok := ApplyValidator(req, v, schemas, src, map[string][]byte{"userSchema": src})
	if !ok || req.Body == nil || len(req.Body.Fields) != 1 {
		t.Fatalf("got ok=%v body=%+v", ok, req.Body)
	}
}

func TestParseChain(t *testing.T) {
	root, src := parseTS(t, `const x = z.string().email().min(3);`)
	calls := findAllByType(root, "call_expression")
	// outermost call is .min(3)
	f := ParseChain(calls[0], src)
	if f.Type != "string" {
		t.Fatalf("type: %+v", f)
	}
	if f.MinLength == nil || *f.MinLength != 3 {
		t.Fatalf("min: %v", f.MinLength)
	}
}

func TestParseSchemaAndObjectArgs(t *testing.T) {
	root, src := parseTS(t, `const s = z.object({ name: z.string() });`)
	calls := findAllByType(root, "call_expression")
	var objCall *sitter.Node
	for _, c := range calls {
		if IsObjectCall(c, src) {
			objCall = c
			break
		}
	}
	if objCall == nil {
		t.Fatal("no object call")
	}
	fields := ParseObjectArgs(objCall, src)
	if len(fields) != 1 || fields[0].Name != "name" {
		t.Fatalf("ParseObjectArgs: %+v", fields)
	}
	schemaFields := ParseSchema(objCall, src)
	if len(schemaFields) != 1 {
		t.Fatalf("ParseSchema: %+v", schemaFields)
	}
}

func TestParseSchema_Nil(t *testing.T) {
	if got := ParseSchema(nil, nil); got != nil {
		t.Fatalf("expected nil, got %+v", got)
	}
}

func TestExtractArgValues(t *testing.T) {
	root, src := parseTS(t, `f("str", 42, ['a','b']);`)
	args := findAllByType(root, "arguments")[0]
	nodes := collectArgNodes(args)
	if got := extractArgValues(nodes[0], src); len(got) != 1 || got[0] != "str" {
		t.Fatalf("string: %v", got)
	}
	if got := extractArgValues(nodes[1], src); len(got) != 1 || got[0] != "42" {
		t.Fatalf("number: %v", got)
	}
	if got := extractArgValues(nodes[2], src); len(got) != 2 {
		t.Fatalf("array: %v", got)
	}
}

func TestParsePair(t *testing.T) {
	root, src := parseTS(t, `const o = { name: z.string() };`)
	pairs := findAllByType(root, "pair")
	f := ParsePair(pairs[0], src)
	if f == nil || f.Name != "name" || f.Type != "string" {
		t.Fatalf("got %+v", f)
	}
}

func TestApplyValidator_Query(t *testing.T) {
	root, src := parseTS(t, `const qSchema = z.object({ limit: z.number() });`)
	schemas := CollectSchemas(root, src)
	req := &scanner.Request{}
	v := ValidatorInfo{Target: "query", SchemaName: "qSchema"}
	ok := ApplyValidator(req, v, schemas, src, map[string][]byte{"qSchema": src})
	if !ok || len(req.Query) != 1 {
		t.Fatalf("got ok=%v query=%+v", ok, req.Query)
	}
}
