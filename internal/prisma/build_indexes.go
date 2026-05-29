//ff:func feature=prisma type=convert control=iteration dimension=1 topic=prisma
//ff:what 블록 @@index([...])로부터 CREATE INDEX 문들 생성
package prisma

import "strings"

// buildIndexes returns CREATE INDEX statements from every @@index([...]) on m.
func buildIndexes(m model, s schema) []string {
	table := tableNameOf(m.name, s)
	out := make([]string, 0, 2)
	for _, a := range m.blockAttrs {
		if !strings.HasPrefix(a, "@@index") {
			continue
		}
		cols := resolveColumns(m.name, bracketList(a), s)
		name := indexName(table, cols)
		out = append(out, "CREATE INDEX "+name+" ON "+table+" ("+strings.Join(cols, ", ")+")")
	}
	return out
}
