//ff:func feature=prisma type=convert control=iteration dimension=1 topic=prisma
//ff:what 단일 model을 컬럼/제약/인덱스를 갖춘 *ddl.Table로 변환
package prisma

import "github.com/park-jun-woo/codistill/internal/ddl"

// buildTable converts one model into a *ddl.Table, skipping relation fields
// (navigation fields are not stored columns).
func buildTable(m model, s schema) *ddl.Table {
	t := &ddl.Table{Name: tableNameOf(m.name, s)}
	for _, f := range m.fields {
		if isRelationField(f, s.models) {
			continue
		}
		t.Columns = append(t.Columns, buildColumn(f, s))
	}
	t.Constraints = buildConstraints(m, s)
	t.Indexes = buildIndexes(m, s)
	return t
}
