//ff:func feature=ddl type=parse control=sequence
//ff:what SQL 문 분류 후 테이블 상태에 적용
package ddl

import "strings"

// applyStatement classifies a SQL statement and updates the tables map.
func applyStatement(tables map[string]*Table, stmt string) {
	clean := stripLeadingComments(stmt)
	if m := reCreateTable.FindStringSubmatch(clean); m != nil {
		applyCreateTable(tables, cleanTableName(m[1]), clean)
		return
	}
	if m := reCreateIndex.FindStringSubmatch(clean); m != nil {
		applyCreateIndex(tables, cleanTableName(m[1]), clean)
		return
	}
	if m := reDropIndex.FindStringSubmatch(clean); m != nil {
		applyDropIndex(tables, cleanTableName(m[1]))
		return
	}
	if m := reDropTable.FindStringSubmatch(clean); m != nil {
		tableName := strings.ToLower(cleanTableName(m[1]))
		delete(tables, tableName)
		return
	}
	if m := reAlterTable.FindStringSubmatch(clean); m != nil {
		applyAlterTable(tables, cleanTableName(m[1]), m[2])
		return
	}
}
