//ff:func feature=prisma type=test control=iteration dimension=1 topic=prisma
//ff:what buildColumn이 enum 기본값은 인용('USER')하고 비-enum 기본값(now/false/0/문자열)은 불변 유지하는지 검증
package prisma

import "testing"

func TestBuildColumnEnumDefault(t *testing.T) {
	s := schema{enums: map[string]bool{"Role": true}}

	cases := []struct {
		name string
		in   field
		want string
	}{
		{"enum default quoted", field{name: "role", baseType: "Role", attrs: []string{"@default(USER)"}}, `"role" "Role" NOT NULL DEFAULT 'USER'`},
		{"enum no default", field{name: "role", baseType: "Role"}, `"role" "Role" NOT NULL`},
		{"now() unchanged", field{name: "createdAt", baseType: "DateTime", attrs: []string{"@default(now())"}}, `"createdAt" timestamp(3) NOT NULL DEFAULT now()`},
		{"false unchanged", field{name: "active", baseType: "Boolean", attrs: []string{"@default(false)"}}, `"active" boolean NOT NULL DEFAULT false`},
		{"numeric unchanged", field{name: "count", baseType: "Int", attrs: []string{"@default(0)"}}, `"count" integer NOT NULL DEFAULT 0`},
		{"string literal unchanged", field{name: "label", baseType: "String", attrs: []string{`@default("x")`}}, `"label" text NOT NULL DEFAULT 'x'`},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := buildColumn(c.in, s).Raw
			if got != c.want {
				t.Errorf("buildColumn(%+v).Raw = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
