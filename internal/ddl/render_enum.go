//ff:func feature=ddl type=render control=iteration dimension=1
//ff:what 단일 enum 타입의 CREATE TYPE ... AS ENUM 출력 (값은 작은따옴표 인용)
package ddl

import "strings"

// renderEnum writes a single CREATE TYPE ... AS ENUM statement.
func renderEnum(sb *strings.Builder, e EnumType) {
	sb.WriteString("CREATE TYPE ")
	sb.WriteString(e.Name)
	sb.WriteString(" AS ENUM (")
	for i, v := range e.Values {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteByte('\'')
		sb.WriteString(v)
		sb.WriteByte('\'')
	}
	sb.WriteString(");\n")
}
