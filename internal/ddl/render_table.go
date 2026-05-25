//ff:func feature=ddl type=render control=iteration dimension=1
//ff:what 단일 테이블의 CREATE TABLE + CREATE INDEX 출력
package ddl

import "strings"

// renderTable writes a single table's CREATE TABLE + CREATE INDEX statements.
func renderTable(sb *strings.Builder, t *Table) {
	sb.WriteString("CREATE TABLE ")
	sb.WriteString(t.Name)
	sb.WriteString(" (\n")

	total := len(t.Columns) + len(t.Constraints)
	idx := 0
	for _, col := range t.Columns {
		sb.WriteString("    ")
		sb.WriteString(cleanLine(col.Raw))
		idx++
		if idx < total {
			sb.WriteByte(',')
		}
		sb.WriteByte('\n')
	}
	for _, con := range t.Constraints {
		sb.WriteString("    ")
		sb.WriteString(cleanLine(con))
		idx++
		if idx < total {
			sb.WriteByte(',')
		}
		sb.WriteByte('\n')
	}

	sb.WriteString(");\n")

	for _, index := range t.Indexes {
		sb.WriteString(cleanLine(index))
		sb.WriteString(";\n")
	}
}
