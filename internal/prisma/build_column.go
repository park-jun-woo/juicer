//ff:func feature=prisma type=convert control=sequence topic=prisma
//ff:what 스칼라 필드를 ddl.Column{Name, Raw}로 변환 (타입/NOT NULL/DEFAULT/@map)
package prisma

import (
	"strings"

	"github.com/park-jun-woo/codistill/internal/ddl"
)

// buildColumn converts a scalar field into a ddl.Column.
func buildColumn(f field) ddl.Column {
	name := columnName(f)
	sqlType := mapType(f)

	def, hasDefault := defaultClause(f)
	sqlType = promoteSerial(sqlType, def, hasDefault)

	var sb strings.Builder
	sb.WriteString(name)
	sb.WriteByte(' ')
	sb.WriteString(sqlType)
	if !f.nullable {
		sb.WriteString(" NOT NULL")
	}
	if hasDefault && def != "" {
		sb.WriteString(" DEFAULT ")
		sb.WriteString(def)
	}
	return ddl.Column{Name: name, Raw: sb.String()}
}
