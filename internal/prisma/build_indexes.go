//ff:func feature=prisma type=convert control=iteration dimension=1 topic=prisma
//ff:what 블록 @@index([...])로부터 CREATE INDEX 문들 생성
package prisma

import "strings"

// buildIndexes returns CREATE INDEX statements from every @@index([...]) on m.
// The index name is built from raw table/columns then quoted once as a whole,
// while the ON clause and column list use the quoted table/columns.
func buildIndexes(m model, s schema) []string {
	rawTable := rawTableName(m.name, s)
	table := tableNameOf(m.name, s)
	out := make([]string, 0, 2)
	for _, a := range m.blockAttrs {
		if !strings.HasPrefix(a, "@@index") {
			continue
		}
		fields := bracketList(a)
		rawCols := resolveRawColumns(m.name, fields, s)
		cols := resolveColumns(m.name, fields, s)
		name := quoteIdent(indexName(rawTable, rawCols))
		out = append(out, "CREATE INDEX "+name+" ON "+table+" ("+strings.Join(cols, ", ")+")")
	}
	return out
}
