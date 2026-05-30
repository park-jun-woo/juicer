//ff:func feature=scan type=test control=iteration dimension=1 topic=dotnet
//ff:what TestCsharpTypeToOpenAPI_Round5 테스트
package dotnet

import "testing"

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
