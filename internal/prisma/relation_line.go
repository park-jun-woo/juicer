//ff:func feature=prisma type=convert control=sequence topic=prisma
//ff:what @relation 인자를 FOREIGN KEY (fk) REFERENCES table (ref) [ON DELETE/UPDATE] 라인으로 변환
package prisma

import "strings"

// relationLine builds one FOREIGN KEY line from a parsed @relation argument.
// Returns "" when the relation has no fields (non-owning side).
func relationLine(m model, f field, rel string, s schema) string {
	fkFields := namedBracketList(rel, "fields")
	refFields := namedBracketList(rel, "references")
	if len(fkFields) == 0 || len(refFields) == 0 {
		return ""
	}
	fkCols := resolveColumns(m.name, fkFields, s)
	refTable := tableNameOf(f.baseType, s)
	refCols := resolveColumns(f.baseType, refFields, s)

	var sb strings.Builder
	sb.WriteString("FOREIGN KEY (")
	sb.WriteString(strings.Join(fkCols, ", "))
	sb.WriteString(") REFERENCES ")
	sb.WriteString(refTable)
	sb.WriteString(" (")
	sb.WriteString(strings.Join(refCols, ", "))
	sb.WriteString(")")
	sb.WriteString(referentialActions(rel))
	return sb.String()
}
