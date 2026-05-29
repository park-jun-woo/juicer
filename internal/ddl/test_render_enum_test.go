//ff:func feature=ddl type=test control=iteration dimension=1
//ff:what renderEnum이 CREATE TYPE ... AS ENUM (인용된 값) 문을 출력하는지 검증
package ddl

import (
	"strings"
	"testing"
)

func TestRenderEnum(t *testing.T) {
	cases := []struct {
		name string
		in   EnumType
		want string
	}{
		{"two values", EnumType{Name: "Role", Values: []string{"USER", "ADMIN"}}, "CREATE TYPE Role AS ENUM ('USER', 'ADMIN');\n"},
		{"single value", EnumType{Name: "Status", Values: []string{"OK"}}, "CREATE TYPE Status AS ENUM ('OK');\n"},
		{"no values", EnumType{Name: "Empty"}, "CREATE TYPE Empty AS ENUM ();\n"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var sb strings.Builder
			renderEnum(&sb, c.in)
			if got := sb.String(); got != c.want {
				t.Errorf("renderEnum(%+v) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
