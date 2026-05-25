//ff:func feature=ddl type=parse control=iteration dimension=2
//ff:what CREATE TABLE 파싱 후 테이블 등록
package ddl

import "strings"

// applyCreateTable parses a CREATE TABLE statement and registers the table.
func applyCreateTable(tables map[string]*Table, name, stmt string) {
	name = strings.ToLower(name)

	// Strip inline comments before extracting body to avoid comment text
	// being treated as column definitions after comma splitting
	cleaned := stripInlineComments(stmt)

	body := extractParenBody(cleaned)
	if body == "" {
		tables[name] = &Table{Name: name}
		return
	}

	t := &Table{Name: name}
	lines := splitTopLevel(body, ',')

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if reConstraintLine.MatchString(line) {
			t.Constraints = append(t.Constraints, line)
		} else {
			colName := extractColumnName(line)
			if colName == "" {
				continue
			}
			t.Columns = append(t.Columns, Column{Name: colName, Raw: line})
		}
	}

	tables[name] = t
}
