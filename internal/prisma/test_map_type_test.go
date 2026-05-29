//ff:func feature=prisma type=test control=iteration dimension=1
//ff:what mapType 타입 매핑 테이블 테스트 (스칼라/@db native/Unsupported/nullable)
package prisma

import "testing"

func TestMapType(t *testing.T) {
	cases := []struct {
		name string
		in   field
		want string
	}{
		{"string scalar", field{baseType: "String"}, "text"},
		{"int scalar", field{baseType: "Int"}, "integer"},
		{"bigint scalar", field{baseType: "BigInt"}, "bigint"},
		{"boolean scalar", field{baseType: "Boolean"}, "boolean"},
		{"datetime scalar", field{baseType: "DateTime"}, "timestamp(3)"},
		{"float scalar", field{baseType: "Float"}, "double precision"},
		{"decimal scalar", field{baseType: "Decimal"}, "numeric"},
		{"json scalar", field{baseType: "Json"}, "jsonb"},
		{"bytes scalar", field{baseType: "Bytes"}, "bytea"},
		{"unknown type passthrough (enum name)", field{baseType: "Role"}, "Role"},

		// @db native takes precedence over the scalar table.
		{"db.VarChar native", field{baseType: "String", attrs: []string{"@db.VarChar(255)"}}, "varchar(255)"},
		{"db.Text native", field{baseType: "String", attrs: []string{"@db.Text"}}, "text"},
		{"db.Uuid native", field{baseType: "String", attrs: []string{"@db.Uuid"}}, "uuid"},

		// Unsupported(...) preserves the inner raw text.
		{"unsupported pgvector", field{baseType: `Unsupported("vector(768)")`}, "vector(768)"},
		{"unsupported escaped quotes", field{baseType: `Unsupported("\"public\".geography")`}, `"public".geography`},

		// nullable (trailing ?) does not affect mapType; baseType already stripped.
		{"nullable string still text", field{baseType: "String", nullable: true}, "text"},
		{"nullable unsupported still raw", field{baseType: `Unsupported("vector(768)")`, nullable: true}, "vector(768)"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := mapType(c.in); got != c.want {
				t.Errorf("mapType(%+v) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
