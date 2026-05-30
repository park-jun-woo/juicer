//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what dotnet 순수 함수 분기 커버리지 보강 테스트 (round5)
package dotnet

import "testing"

func TestMethodAllowsBody_Round5(t *testing.T) {
	for _, m := range []string{"POST", "PUT", "PATCH"} {
		if !methodAllowsBody(m) {
			t.Errorf("%s should allow body", m)
		}
	}
	for _, m := range []string{"GET", "DELETE", "HEAD", ""} {
		if methodAllowsBody(m) {
			t.Errorf("%s should not allow body", m)
		}
	}
}

func TestUnquoteCSharp_Round5(t *testing.T) {
	if got := unquoteCSharp(`"hi"`); got != "hi" {
		t.Errorf("quoted: got %q", got)
	}
	if got := unquoteCSharp("x"); got != "x" {
		t.Errorf("short: got %q", got)
	}
	if got := unquoteCSharp("ab"); got != "ab" {
		t.Errorf("no-quotes: got %q", got)
	}
	if got := unquoteCSharp(""); got != "" {
		t.Errorf("empty: got %q", got)
	}
}

func TestCsharpTypeToOpenAPIType_Round5(t *testing.T) {
	if got := csharpTypeToOpenAPIType("int"); got != "integer" {
		t.Errorf("int: got %q", got)
	}
	// void -> empty Type -> "string" fallback
	if got := csharpTypeToOpenAPIType("void"); got != "string" {
		t.Errorf("void fallback: got %q", got)
	}
}

func TestCsharpTypeToOpenAPI_Round5(t *testing.T) {
	cases := []struct {
		in     string
		typ    string
		format string
		items  string
	}{
		{"string", "string", "", ""},
		{" int? ", "integer", "int32", ""},
		{"Int32", "integer", "int32", ""},
		{"long", "integer", "int64", ""},
		{"Int64", "integer", "int64", ""},
		{"short", "integer", "int32", ""},
		{"Int16", "integer", "int32", ""},
		{"float", "number", "float", ""},
		{"Single", "number", "float", ""},
		{"double", "number", "double", ""},
		{"Double", "number", "double", ""},
		{"decimal", "number", "", ""},
		{"Decimal", "number", "", ""},
		{"bool", "boolean", "", ""},
		{"Boolean", "boolean", "", ""},
		{"DateTime", "string", "date-time", ""},
		{"DateTimeOffset", "string", "date-time", ""},
		{"DateOnly", "string", "date", ""},
		{"Guid", "string", "uuid", ""},
		{"IFormFile", "string", "binary", ""},
		{"byte[]", "string", "byte", ""},
		{"void", "", "", ""},
		{"List<Foo>", "array", "", "Foo"},
		{"IEnumerable<Bar>", "array", "", "Bar"},
		{"IList<A>", "array", "", "A"},
		{"ICollection<B>", "array", "", "B"},
		{"Dictionary<string,int>", "object", "", ""},
		{"IDictionary<string,int>", "object", "", ""},
		{"Widget[]", "array", "", "Widget"},
		{"SomeDTO", "object", "", ""},
	}
	for _, c := range cases {
		oa := csharpTypeToOpenAPI(c.in)
		if oa.Type != c.typ || oa.Format != c.format || oa.Items != c.items {
			t.Errorf("%q: got %+v, want type=%q format=%q items=%q", c.in, oa, c.typ, c.format, c.items)
		}
	}
}

func TestExtractGenericInner_Round5(t *testing.T) {
	if got := extractGenericInner("List<Foo>"); got != "Foo" {
		t.Errorf("got %q", got)
	}
	if got := extractGenericInner("NoGeneric"); got != "NoGeneric" {
		t.Errorf("no-open: got %q", got)
	}
	if got := extractGenericInner("Broken<"); got != "Broken<" {
		t.Errorf("no-close: got %q", got)
	}
	if got := extractGenericInner(">Bad<"); got != ">Bad<" {
		t.Errorf("end<=start: got %q", got)
	}
}

func TestUnwrapReturnType_Round5(t *testing.T) {
	cases := []struct {
		in    string
		inner string
		array bool
	}{
		{"ActionResult<Foo>", "Foo", false},
		{"ActionResult<List<Foo>>", "Foo", true},
		{"Task<Bar>", "Bar", false},
		{"List<X>", "X", true},
		{"IEnumerable<Y>", "Y", true},
		{"IList<Z>", "Z", true},
		{"ICollection<W>", "W", true},
		{"Thing[]", "Thing", true},
		{"Plain", "Plain", false},
	}
	for _, c := range cases {
		inner, arr := unwrapReturnType(c.in)
		if inner != c.inner || arr != c.array {
			t.Errorf("%q: got (%q,%v), want (%q,%v)", c.in, inner, arr, c.inner, c.array)
		}
	}
}

func TestStripRouteConstraints_Round5(t *testing.T) {
	if got := stripRouteConstraints("api/{id:int}"); got != "api/{id}" {
		t.Errorf("constraint: got %q", got)
	}
	if got := stripRouteConstraints("api/{*slug}"); got != "api/{slug}" {
		t.Errorf("catchall: got %q", got)
	}
	if got := stripRouteConstraints("api/plain"); got != "api/plain" {
		t.Errorf("plain: got %q", got)
	}
}

func TestIsDIType_Round5(t *testing.T) {
	if !isDIType("ILogger<Foo>") {
		t.Error("ILogger should be DI")
	}
	if !isDIType("AppDbContext") {
		t.Error("DbContext should be DI")
	}
	if !isDIType("IUserService") {
		t.Error("IUserService should be DI")
	}
	if !isDIType("IOrderRepository") {
		t.Error("IOrderRepository should be DI")
	}
	if isDIType("IUserThing") {
		t.Error("IUserThing should not be DI")
	}
	if isDIType("string") {
		t.Error("string should not be DI")
	}
	if isDIType("I") {
		t.Error("single I should not be DI")
	}
	if isDIType("iLower") {
		t.Error("lowercase second char should not be DI")
	}
}
