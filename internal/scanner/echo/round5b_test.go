//ff:func feature=scan type=test control=sequence topic=echo
//ff:what round5 struct/type 해석 + 파라미터 핸들러 테스트 (echo)
package echo

import (
	"go/ast"
	"go/types"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// namedType returns the *types.Named for the package-level var `name`.
func namedType(t *testing.T, info *types.Info, name string) types.Type {
	t.Helper()
	for id, obj := range info.Defs {
		if obj != nil && id.Name == name {
			return obj.Type()
		}
	}
	t.Fatalf("no type for %s", name)
	return nil
}

const dtoSrc = `package m
type Inner struct {
	Code int ` + "`json:\"code\"`" + `
}
type Base struct {
	ID int ` + "`json:\"id\"`" + `
}
type UserDto struct {
	Base
	Name   string ` + "`json:\"name\"`" + `
	Hidden string ` + "`json:\"-\"`" + `
	Nested Inner  ` + "`json:\"nested\"`" + `
}
var U UserDto
var L []UserDto
`

func TestResolveType_Round5(t *testing.T) {
	_, info := checkSrc(t, dtoSrc)
	typ := namedType(t, info, "U")
	name, fields := resolveType(typ)
	if name != "UserDto" {
		t.Fatalf("name: %q", name)
	}
	if len(fields) == 0 {
		t.Fatalf("expected fields, got %+v", fields)
	}
	// json:"-" must be excluded
	for _, f := range fields {
		if f.Name == "Hidden" {
			t.Errorf("Hidden should be excluded")
		}
	}

	// slice type -> []UserDto prefix
	lname, _ := resolveType(namedType(t, info, "L"))
	if lname != "[]UserDto" {
		t.Fatalf("slice name: %q", lname)
	}
}

func TestExtractFields_And_BuildField_Round5(t *testing.T) {
	_, info := checkSrc(t, dtoSrc)
	typ := namedType(t, info, "U")
	named := typ.(*types.Named)
	st := named.Underlying().(*types.Struct)
	fields := extractFields(st, map[string]bool{})
	// embedded Base.ID should be flattened in
	var hasID, hasName bool
	for _, f := range fields {
		if f.JSON == "id" {
			hasID = true
		}
		if f.JSON == "name" {
			hasName = true
		}
	}
	if !hasID {
		t.Errorf("embedded id not flattened: %+v", fields)
	}
	if !hasName {
		t.Errorf("name missing: %+v", fields)
	}

	// buildField with json:"-" returns nil
	var hiddenVar *types.Var
	for i := 0; i < st.NumFields(); i++ {
		if st.Field(i).Name() == "Hidden" {
			hiddenVar = st.Field(i)
		}
	}
	if hiddenVar == nil {
		t.Fatal("no hidden field")
	}
	if buildField(hiddenVar, `json:"-"`, map[string]bool{}) != nil {
		t.Error("json:\"-\" field should be nil")
	}
}

func TestResolveNestedFields_And_Embedded_Round5(t *testing.T) {
	_, info := checkSrc(t, dtoSrc)
	typ := namedType(t, info, "U")
	nested := resolveNestedFields(typ, map[string]bool{})
	if len(nested) == 0 {
		t.Fatalf("expected nested fields, got %+v", nested)
	}

	named := typ.(*types.Named)
	emb := resolveEmbedded(named, map[string]bool{})
	if len(emb) == 0 {
		t.Fatalf("expected embedded resolution, got %+v", emb)
	}
}

func TestHandlePathParam_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := callWithStringArg(t, `"id"`)
	handlePathParam(ep, call)
	if ep.Request == nil || len(ep.Request.PathParams) != 1 || ep.Request.PathParams[0].Name != "id" {
		t.Fatalf("path params: %+v", ep.Request)
	}
	// duplicate should be ignored
	handlePathParam(ep, call)
	if len(ep.Request.PathParams) != 1 {
		t.Fatalf("duplicate added: %+v", ep.Request.PathParams)
	}
	// empty arg -> no-op
	ep2 := &scanner.Endpoint{}
	handlePathParam(ep2, &ast.CallExpr{})
	if ep2.Request != nil {
		t.Fatalf("expected nil request for no args")
	}
}

func TestHandleQueryParam_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := callWithStringArg(t, `"page"`)
	handleQueryParam(ep, call, "GET")
	if ep.Request == nil || len(ep.Request.Query) != 1 || ep.Request.Query[0].Name != "page" {
		t.Fatalf("query: %+v", ep.Request)
	}
	// duplicate ignored
	handleQueryParam(ep, call, "GET")
	if len(ep.Request.Query) != 1 {
		t.Fatalf("duplicate query: %+v", ep.Request.Query)
	}
}

// callWithStringArg builds an *ast.CallExpr whose first arg is the given literal.
func callWithStringArg(t *testing.T, lit string) *ast.CallExpr {
	t.Helper()
	expr := parseExpr(t, "f("+lit+")")
	call, ok := expr.(*ast.CallExpr)
	if !ok {
		t.Fatalf("not a call expr: %s", lit)
	}
	return call
}
